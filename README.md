# pure-netcopy

Simple file transfer on TCP basis to migrate data from old NAS.

## Story 

I wanted to transfer my files, mostly media files, from my old NAS Qnap TS-410 to my new diy NAS. The two NAS are connected over 1Gb network and I thought that it would be straightforward, but as I tried to transfer the files via SFTP the transfer speed reached only 36Mb/s. After checking the network and a couple of other transfer methods I realized that the main issue is the performance of my old NAS (old ARM CPU). As soon as any conversion, encryption or compression happens the performance drops. I verified it with netcat to stream a file from the old NAS to the new NAS and I doubled the speed to 76 Mb/s. I was considering TFTP to transfer the files, but I was concerned with UDP. I tried to hash the inputfile, then transfer it to the server and there check the hash again, but hashing the file even with simple crc took forever because of the limited capacity of the ARM CPU. At this time I stopped looking and started this little project. There may be other tools out there who do similar things, but I thought this would be the faster method.