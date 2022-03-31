// started with https://www.educative.io/edpresso/how-to-implement-tcp-sockets-in-c
// and the example in https://www.man7.org/linux/man-pages/man3/getaddrinfo.3.html

#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>

#include <netdb.h> // For getaddrinfo
#include <unistd.h> // for close
#include <stdlib.h> // for exit
                    //
#include <assert.h>
#include <stdbool.h>
                   
#define PORT 8301
#define BUFFER_SIZE 1024

typedef struct Response {
   char *body;
   char *statusCode;
   char *statusMessage;
   bool hasBody;
} response;

int openConnection ();
char* sendRequest (char* request);
response* fetchResponse(char* request);

/* helper functions */
char** splitStrings (char* line, char* delimeter);
char*  trimTrailing (char *s);

void printResponse(response* res) {
    printf("Respone:\n");
    printf("Status code: %s\n", res->statusCode);
    printf("Status message: %s\n", res->statusMessage);
    if (res->hasBody) {
        printf("Body: %s\n", res->body);
    }
}

int main(void)
{

    char* uuid = "Kien";
    char* content = "a random one";

    char **temp1, **temp2;
    char *temp4, *temp5, *temp6;
    char *expectedGet2;

    response *resGet1;
    response *resPost;
    response *resGet2;
    response *resDelete;
    response *resPut;

    char *requestGet;
    char *requestPost;
    char *requestDelete;
    char *requestPut;

    char *id;

    //Get all memos
    printf("\n\n>>> Get all memos\n");
    requestGet = "GET /api/memo HTTP/1.1\r\nHost: 127.0.0.1:8301\r\nCookie: uuid=%s\n\r\n";
    printf("> Request:\n%s\n", requestGet);
    resGet1 = fetchResponse(requestGet);
    printResponse(resGet1);

    assert(strcmp(resGet1->statusCode, "200") == 0);

    //Post one memo
    printf("\n\n>>> Insert memo\n");
    requestPost = "POST /api/memo/ HTTP/1.1\nHost: 127.0.0.1:8301\nContent-Type: application/json\nCookie: uuid=Kien\nContent-Length: 39\n\n{\"content\": \"Rand domeas\"}";
    printf("> Request:\n%s\n", requestPost);
    resPost = fetchResponse(requestPost);
    printResponse(resPost);

    assert(strcmp(resPost->statusCode, "200") == 0);

    //Get id
    temp1 = splitStrings(resPost->body, ","); //Split into pair
    temp4 = temp1[0]; //First pair is id
    printf("%s\n", temp4);
    temp2 = splitStrings(temp4, ":"); //Split into key - value
    temp5 = temp2[1]; 
    printf("-%s\n", temp5);
    
    id = temp5;

    
    //Get and compare
    printf("\n\n>>> Get all memos\n");
    printf("> Request:\n%s\n", requestGet);
    resGet2 = fetchResponse(requestGet);
    printResponse(resGet2);

    assert(strcmp(resGet2->statusCode, "200") == 0);

    //Compare
    /* Concanate get1 body and insert body -> get 2 body */
    temp1 = splitStrings(resGet1->body, "]");
    temp4 = temp1[0];

    temp5 = strcat(temp4, ",");
    temp6 = strcat(temp5, resPost->body);

    expectedGet2 = strcat(temp6, "]");

    printf("Expected get 2:\n%s\n", expectedGet2);

    assert(strcmp(expectedGet2, resGet2->body) == 0);


    //Delete
    printf("\n\n>>> Delete memo\n");
    requestDelete = (char *) malloc(sizeof(char) * BUFFER_SIZE);

    sprintf(requestDelete, "DELETE /api/memo/%s HTTP/1.1\nHost: 127.0.0.1:8301\nCookie: uuid=%s", id, uuid);
    printf("> Request:\n%s\n", requestDelete);

    resDelete = fetchResponse(requestDelete);
    printResponse(resDelete);

    assert(strcmp(resDelete->statusCode, "200") == 0);

    //Put to update -> expect to get 400 Bad Request response
    printf("\n\n>>> Put/Update memo to test deletion\n");
    requestPut = (char *) malloc(sizeof(char) * BUFFER_SIZE);

    sprintf(requestPut, "PUT /api/memo/ HTTP/1.1\nHost: 127.0.0.1:8301\nContent-Type: application/json\nCookie: uuid=%s\n\n{\"id\": %s,\"content\": \"Random content\"}", id, uuid);
    printf("> Request:\n%s\n", requestPut);
    resPut = fetchResponse(requestPut);
    printResponse(resPut);

    assert(strcmp(resPut->statusCode, "404") == 0);

    free(requestPut);
    free(requestDelete);
}

response* fetchResponse (char *request) {
    response *res;
    char *server_message;
    char **temp;
    char *tempLine;
    char **responseLineTokens;
    int tempSize;
    int i;
    int sectionBreak;
    char *responseLine;

    res = (response*) malloc(sizeof(response));

    server_message = sendRequest(request);

    temp = splitStrings(server_message, "\n");

    res->hasBody = false;
    i = 0;
    tempSize = 0;

    while (temp[i] != NULL) {
        tempLine = temp[i];
        if (strlen(tempLine) == 0) {
            sectionBreak = i;
        }

        i ++;
        tempSize ++;
    }

    responseLine = temp[0];

    if (sectionBreak < tempSize-1) {
        res->hasBody = true;
        res->body = temp[sectionBreak+1];
    }

    responseLineTokens = splitStrings(responseLine, " ");

    res->statusCode = responseLineTokens[1];
    res->statusMessage = responseLineTokens[2];

    return res;

}

int openConnection () {
    int socket_desc;
    struct sockaddr_in server_addr;
    char address[100];

    struct addrinfo *result;
    
    
    // Create socket:
    socket_desc = socket(AF_INET, SOCK_STREAM, 0);
    
    if(socket_desc < 0){
        printf("Unable to create socket\n");
        return -1;
    }
    

    struct addrinfo hints;
    memset (&hints, 0, sizeof (hints));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_flags |= AI_CANONNAME;
    
    // get the ip of the page we want to scrape
    int out = getaddrinfo ("localhost", NULL, &hints, &result);
    // fail gracefully
    if (out != 0) {
        fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(out));
        exit(EXIT_FAILURE);
    }
    
    // ai_addr is a struct sockaddr
    // so, we can just use that sin_addr
    struct sockaddr_in *serverDetails =  (struct sockaddr_in *)result->ai_addr;
    
    // Set port and IP the same as server-side:
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(PORT);
    //server_addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    server_addr.sin_addr = serverDetails->sin_addr;
    
     // converts to octets
    inet_ntop (server_addr.sin_family, &server_addr.sin_addr, address, 100);
    // Send connection request to server:
    if(connect(socket_desc, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0){
        printf("Unable to connect\n");
        exit(EXIT_FAILURE);
    }

    return socket_desc;
} 

char* sendRequest (char* request) {

    char* server_message;
    char client_message[2000];
    
    int socket_desc;

    // Clean buffers:
    memset(client_message,'\0',sizeof(client_message));

    server_message = (char*) malloc(sizeof(char)*BUFFER_SIZE);

    socket_desc = openConnection ();
  
    // Send the message to server:
    if(send(socket_desc, request, strlen(request), 0) < 0){
        printf("Unable to send message\n");
        return NULL;
    }
    
    // Receive the server's response:
    if(recv(socket_desc, server_message, sizeof(char)*4096, 0) < 0){
        printf("Error while receiving server's msg\n");
        return NULL;
    }
    
    printf("Message: %s", server_message);
    
    // Close the socket:
    close(socket_desc);
    
    return server_message;
}


char** splitStrings (char* givenLine, char *delimeter) {
    char **args;
    char * token ;
    int n_spaces;
    char *line;

    args = NULL;
    line = strdup(givenLine);
    token = strsep(&line, delimeter);
    n_spaces = 0;

    while (token) {
        n_spaces ++;
        args = realloc (args, sizeof (char *) * n_spaces);
        if (args == 0) {
            printf("We have errors with reallocation\n");
            exit(EXIT_FAILURE);
        }

        args[n_spaces-1] = token;

        token = strsep (&line, delimeter);
    }

    n_spaces ++;
    args = realloc (args, sizeof (char *) * n_spaces);

    args[n_spaces-1] = NULL;

    return args;
}


char* trimTrailing(char *givenS) {
    char * s;
	int i;
    size_t j;
    char *newS;

    s = strdup (givenS);
    /* trim postfix */
    i = strlen(s)-1;
	while(i >= 0) {
	  if(s[i]==' '||s[i]=='\t')
	  i--;
	  else
	   break;
	}
	s[i+1]='\0';

    /* trim prefix */
    j = 0;

    while (j < strlen(s) && (s[j] == ' ' || s[j] == '\t')) {
        j++;
    }

    if ( j < strlen(s)) {
        newS = &s[j];
        return newS;
    } else {
        return NULL;
    }
}
