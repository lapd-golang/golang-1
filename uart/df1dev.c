#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <termios.h>
#include <fcntl.h>
#include <errno.h>
#include <ctype.h>

// serial
#define BUF_SIZE	512	

static struct termios save_termios;
static int  ttysavefd = -1;
static enum {RESET , RAW , CBREAK} ttystate = RESET;

int speed_arr[] = { B115200, B57600, B38400, B19200, B9600, B4800, B2400, B1200, B300};
int name_arr[] = {115200, 57600, 38400, 19200, 9600, 4800, 2400, 1200, 300};

unsigned char crc8(unsigned char arr[], int len)
{
	int i = 0; 
	unsigned char crc = 0;

	for(i = 0; i < len; i++){
		crc ^= arr[i];
	}

	return crc;
}

unsigned short calc_crc (unsigned short crc, unsigned short buffer) 
{
	unsigned short temp1, y;
	temp1 = crc ^ buffer;
	crc = (crc & 0xff00) | (temp1 & 0xff);
	for (y = 0; y < 8; y++){
	   if (crc & 1){	  
			crc = crc >> 1;
			crc ^= 0xa001;
		} else
			crc = crc >> 1;
	}
	return crc;
}

unsigned short compute_crc(unsigned char buf[], int len) 
{
	int x;
	unsigned short crc = 0;
	for (x = 0; x < len; x++) { 
		crc = calc_crc(crc, buf[x]);
	}
	crc = calc_crc(crc, 0x03);
	return (crc);
}

void set_speed(int fd, int speed)
{
        int   i;
        int   status;
        struct termios   Opt;
        tcgetattr(fd, &Opt);
        for ( i= 0;  i < sizeof(speed_arr) / sizeof(int);  i++) {
                if  (speed == name_arr[i]) {
                        tcflush(fd, TCIOFLUSH);
                        cfsetispeed(&Opt, speed_arr[i]);
                        cfsetospeed(&Opt, speed_arr[i]);
                        status = tcsetattr(fd, TCSANOW, &Opt);
                        if  (status != 0) {
                                perror("tcsetattr fd");
                                return;
                        }
                        tcflush(fd,TCIOFLUSH);
                }
        }

        return ;
}

int set_Parity(int fd,int databits,int stopbits,int parity)
{
	struct termios options;
	if  ( tcgetattr( fd,&options)  !=  0) {
		perror("SetupSerial 1");
		return 1;
	}
	options.c_cflag &= ~CSIZE;
	switch (databits)
	{
		case 7:
			options.c_cflag |= CS7;
			break;
		case 8:
			options.c_cflag |= CS8;
			break;
		default:
			fprintf(stderr,"Unsupported data size\n"); return 1;
	}
	switch (parity)
	{
		case 'n':
		case 'N':
		case 0:
			options.c_cflag &= ~PARENB;   /* Clear parity enable */
			options.c_iflag &= ~INPCK;     /* Enable parity checking */
			break;
		case 'o':
		case 'O':
		case 2:
			options.c_cflag |= (PARODD | PARENB);
			options.c_iflag |= INPCK;             /* Disnable parity checking */
			break;
		case 'e':
		case 'E':
		case 1:
			options.c_cflag |= PARENB;     /* Enable parity */
			options.c_cflag &= ~PARODD;
			options.c_iflag |= INPCK;       /* Disnable parity checking */
			break;
		case 'S':
		case 's':  /*as no parity*/
			options.c_cflag &= ~PARENB;
			options.c_cflag &= ~CSTOPB;break;
		default:
			fprintf(stderr,"Unsupported parity\n");
			return 1;
	}
	switch (stopbits)
	{
		case 1:
			options.c_cflag &= ~CSTOPB;
			break;
		case 2:
			options.c_cflag |= CSTOPB;
			break;
		default:
			fprintf(stderr,"Unsupported stop bits\n");
			return 1;
	}
	/* Set input parity option */
	if (parity != 'n')
		options.c_iflag |= INPCK;
	tcflush(fd,TCIFLUSH);
	options.c_cc[VTIME] = 0;
	options.c_cc[VMIN] = 1; /* Update the options and do it NOW */
	if (tcsetattr(fd,TCSANOW,&options) != 0)
	{
		perror("SetupSerial 3");
		return 1;
	}
	return 0;
}

int tty_raw(int fd)
{
        int err;
        struct termios buf;

        if(ttystate != RESET){
                errno = EINVAL;
                return -1;
        }
        if(tcgetattr(fd, &buf)<0)
                return -1;
        save_termios = buf;

        buf.c_lflag &= ~(ECHO | ICANON | IEXTEN | ISIG);

        buf.c_iflag &= ~(BRKINT | ICRNL | INPCK | ISTRIP | IXON);

        buf.c_cflag &= ~(CSIZE | PARENB);

        buf.c_cflag |= CS8;

        buf.c_oflag &= ~(OPOST);
        buf.c_cc[VMIN] = 1;
        buf.c_cc[VTIME] = 0;
        if(tcsetattr(fd, TCSAFLUSH, &buf)<0)
                return -1;

        if(tcgetattr(fd, &buf)<0){
                err = errno;
                tcsetattr(fd, TCSAFLUSH, &save_termios);
                errno = err;
                return -1;
        }

        if((buf.c_lflag & (ECHO | ICANON | IEXTEN | ISIG)) || (buf.c_iflag & (BRKINT | ICRNL | INPCK | ISTRIP | IXON )) || (buf.c_cflag & ( CSIZE | PARENB | CS8)) != CS8 || (buf.c_oflag & OPOST) || buf.c_cc[VMIN]!=1 || buf.c_cc[VTIME]!=0){
                tcsetattr(fd, TCSAFLUSH, &save_termios);
                errno = EINVAL;
                return -1;

        }

        ttystate = RAW;
        ttysavefd = fd;
        return 0;

}

int main(int argc, char *argv[])
{
	int fd = 0;
	unsigned char tx_buf[BUF_SIZE], rx_buf[BUF_SIZE];
	int speed = 0, len = 0;
	char databits, stopbits, parity;
	int i = 0, count = 0, num = 0, sum = 0;
	unsigned short crc = 0;
/*
	unsigned char buf[]={0x01, 0x00, 0x0F, 0x00, 0x07, 0x43, 0xA2, 0x02, 0x07, 0x89, 0x65, 0x00};
	crc = compute_crc(buf, 12);
	printf("0x%X, 0x%X\n", crc&0xFF, (crc>>8 &0xFF));
	exit(0);
*/
	
	printf("Usage: %s /dev/ttyxxx mode speed databits parity stopbits\n", argv[0]);
	printf("mode: specify the mode , full or half, not using, just for the same.\n");
	printf("speed: specify the bps, 115200, 57600, 9600, 4800, 2400...\n");
	printf("parity: 0:none, 1:even, 2:odd\n");

	if(7 != argc){
		printf("CMD error! Please repeat!\n");	
		exit(-1);
	}

	speed = atoi(argv[3]);
	databits = argv[4][0] - '0';
	parity = argv[5][0] - '0';
	stopbits = argv[6][0] - '0';

	fd = open(argv[1], O_RDWR);
	if(fd == -1){
		perror("open ttyUSBx");
		exit(1);
	}
#if 1
	tty_raw(fd);

	set_speed(fd, speed);
	set_Parity(fd, databits, stopbits, parity);
#endif

	while(1){
		sum = 0;
		do {
			len = read(fd, rx_buf+sum, sizeof(rx_buf));
			if(len < 0){
				continue;
			}
			sum += len;
//		} while(0);
//		} while((sum<10 || (rx_buf[sum-4] != 0x10 && rx_buf[sum-3] != 0x03))&&(sum < 30));
		} while((sum<4 || (rx_buf[sum-4] != 0x10 && rx_buf[sum-3] != 0x03)));

		printf("receive %d [hex] data:", sum);
		for(i = 0; i < sum; i++){
			if(i % 10 == 0)printf("\n");
			printf("%.2X ", rx_buf[i]&0xFF);
		}
		printf("\n");
	
		if((rx_buf[0] == 0x10) && (rx_buf[1] == 0x02)){
			tx_buf[0]=0x10;
			tx_buf[1]=0x06;
			write(fd, tx_buf, 2);
			sync();

			tx_buf[0] = 0x10;
			tx_buf[1] = 0x02;
			tx_buf[2] = rx_buf[3];
			tx_buf[3] = rx_buf[2];
			if(rx_buf[4] == 0x0F){
				tx_buf[4] = 0x4F;
			} else {
				printf("Not supported CMD: 0x%.2X\n", rx_buf[4]);
			}
			tx_buf[5] = rx_buf[5];	
			tx_buf[6] = rx_buf[6];
			tx_buf[7] = rx_buf[7];
			
			if(rx_buf[8] == 0xA2){ // FNC: A2, to read
				count = rx_buf[9];	
				for(i=0; i <count; i++){ //  the count of bytes to read
					tx_buf[8+i] = 0x01 + i;
				}
				num = 6 + count;
			} else if(rx_buf[8] == 0xAA){ // FNC: AA, to write
				num = 6;
			}
			crc = compute_crc(&tx_buf[2], num);
			tx_buf[num + 2] = 0x10;
			tx_buf[num + 3] = 0x03;
			tx_buf[num + 4] = crc & 0xFF;
			tx_buf[num + 5] = (crc>>8) & 0xFF;
			num = num + 6;
		} else {
			printf("Repeat send...\n");
		}
		write(fd, tx_buf, num);
		sync();

		printf("send %d [hex] data:\n0x10 0x06", num+2);
		for(i = 0; i < num; i++){
			if(i % 10 == 0)printf("\n");
			printf("%.2X ", tx_buf[i]&0xFF);
		}
		printf("\n");
		printf("-----------------\n");
		
	}

	close(fd);
	return 0;
}

