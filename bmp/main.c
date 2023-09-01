#include <stdio.h>

char Wall[256*256*3] = {0};

int main(){
	printf("Test\n");
	loadBMP(Wall, "Tree.bmp");
	
	FILE *fptr;
	fptr = fopen("bmp.txt", "w");
	
	
	//char Wall_str[256*256*16*3];
	for(int i = 0; i < 256*10*3; i++){ // 256*256*3
		//Wall_str[i] = Wall[i];
		fprintf(fptr, "0x%x, ", Wall[i]);
		printf("%x, ", Wall[i]);
		
	}
	
	//fprintf(fptr, "%s", Wall_str);
	
	fclose(fptr);
}


void loadBMP(char *Wall, char *name){
    FILE *f;

    char BM[3] = {0};
    int sizeFile = 0;
    int offset = 0;
    int width = 0;
    int height = 0;
    short int colors = 0;
    int compress = 0;
    char pixel[3] = {0};
    char buf[30] = {0};

    f = fopen(name, "rb"); // Îòêðûâàåò ôàéë â áèíàðíîì ðåæèìå
            fread(&BM, 1, 2, f);
            fread(&sizeFile, 1, 4, f);
            fseek(f, 0x0A, SEEK_SET);   // Çàäàåò ïîçèöèþ êóðñîðà îò íà÷àëà ôàéëà
            fread(&offset, 1, 4, f);/////////////
            fseek(f, 0x12, SEEK_SET);
            fread(&width, 1, 4, f);//////////////
            fread(&height, 1, 4, f);/////////////
            fseek(f, 0x1C, SEEK_SET);
            fread(&colors, 1, 2, f);
            fread(&compress, 1, 4, f);

        fseek(f, offset, SEEK_SET);
        for(int i = 0; i < height; i++){
            for(int j = 0; j < width; j++){
                fread(&pixel, 1, 3, f);
                Wall[i*width*3+j*3] = (char)pixel[2];
                Wall[i*width*3+j*3+1] = (char)pixel[1];
                Wall[i*width*3+j*3+2] = (char)pixel[0];
            }
        if(width%4 != 0) fread(&buf, 1, ((width/4 + 1)*4) - width, f); // Êîëè÷åñòâî ïèêñåëåé â ñòðîêå äîëæíî áûòü êðàòíî 4
        }
    fclose(f);
}
