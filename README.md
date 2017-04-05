# bs-directory-transverser

The bs-directory-transverser is a directory script includes modifier for brightscript applications. The application was developed avoid scenarios where developers have to manually add every file as a script include to the top node. The application will search through every xml file looking for the following tags `<!-- SCRIPT-INCLUDE uri="<DIRECTORY-TO-INCLUDE>" -->`. When it finds a tag it will transverse through all the files in `DIRECTORY-TO-INCLUDE` and adds all the brs files it finds to the xml as script includes. Multiple directory tags can be placed in each xml file. 

The script is designed to be run at build.

## Dev Dependencies

1. GO 1.7 (Installation process: https://nats.io/documentation/tutorials/go-install/)
2. Docker (Installation process: https://getcarina.com/docs/tutorials/docker-install-mac/)

## Docker Process

1. To build a new image run the command:

 ```docker build -t tcrands/bs-directory-transverser .```

2. To run the image cd to the folder you wish to run it on and run the command:

 ```docker run -v $(pwd):/source tcrands/go-directory-transverser```

3. To tag a new image run `docker images` to get the image id. The run the command:

```docker tag <IMAGEID> tcrands/bs-directory-transvser:<TAG>```

4. To login and push the image run `docker login` and he run the command:

```docker push tcrands/bs-directory-transverser```

5. To run the image in bash mode run the command: (Running the image will pull it from the repo)

```docker run -v $(pwd):/source -it tcrands/go-directory-transverser /bin/bash```
