# Gossip-Cache

 Gossip-Cache is a project that focuses on implementing a cache synchronization mechanism for distributed systems. It utilizes the [Gossip Protocol](https://en.wikipedia.org/wiki/Gossip_protocol) to ensure that all nodes within a network maintain a synchronized cache state, even when faced with the challenges of [Eventual Consistency](https://en.wikipedia.org/wiki/Eventual_consistency) . This approach helps in enhancing the overall performance and reliability of distributed systems by ensuring that cached data remains consistent across all nodes, thereby reducing latency and improving data access times.
<br><br>

<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
  </details>

<!--About the Project-->
  
## About The Project

### Demo
https://github.com/Uttkarsh-raj/Gossip-Cache/assets/106571927/328e2194-1468-430a-b933-d6c4b05c506b

<br>
The above demo illustrates the key features of Gossip-Cache. It begins by simulating the addition of nodes, each initialized with a cache, to the server. Two clients are then introduced to showcase how their caches are shared. The server serves as a gateway for the nodes, facilitating their entry into the network. This demonstration effectively highlights the concept of eventual consistency, as the data may not be immediately available to all nodes but is eventually synchronized across the network. The synchronization time is directly influenced by the number of nodes in the network.

### Key Concepts
1. **Distributed Systems** : The project is fundamentally a distributed system since it involves multiple nodes (clients) that coordinate and communicate to achieve a common goal.
3. **Gossip Protocol** : Gossip protocols are used for information dissemination in distributed systems, ensuring eventual consistency by spreading information in a manner similar to how gossip spreads in social networks.
4. **Eventual Consistency** : The system aims for eventual consistency, meaning that all nodes will converge to the same state over time, even though they may be temporarily inconsistent.
5. **Distributed Caching** : The use of caches in each node and the management of these caches using TTL (time-to-live) values and exchange protocols is a classic problem in distributed caching.

### Key Features
1. **Node Registration**: Each node connecting to the server is registered and stored, creating a robust network map.
2. **Neighbor Discovery**: New nodes dynamically discover and record their neighbors, fostering a resilient network structure.
3. **Cache Exchange**: Nodes use the gossip protocol to exchange cache information, promoting widespread data consistency.
4. **TTL Management**: Cache items are managed with a time-to-live (TTL) mechanism, ensuring outdated entries are purged.
5. **Eventual Consistency**: Designed for environments where eventual consistency is acceptable, Gossip-Cache excels in maintaining a balanced cache state.




### Built With
<br><br>

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build and execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.



## Installation 

1. Create an image from the docker file:
   
   ```
   docker build -t gossipcache .
   ```
3. Run this on your terminal (needs docker to be preinstalled):
   
   ```
   docker run -p 3000:3000 -it gossipcache
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

- **Get "/"**
  * Connects to the gateway server and stay connected.
  * Response as :
    ```
    {
      "data": null,
      "message": "Successfully connected to server.",
      "success": true
    }
    ```
      
- **Get "/get/:key"**
  * Response as :
    ```
    {
      "data": {
      "key": "test",
      "value": "data",
      "ttl": 1718998692919380103
      },
      "message": "Successfully data retrieved",
      "success": true
    }
    ```
      
- **POST "/set"**
  * Request as (only etherium/bitcoin are possible values currently):
    - key : The unique string which will be used to fetch this data.
    - value : Can be anything (string ,int, array...) whose value is to be stored.
    - ttl : TTL is the time after which the system should delete this data (in seconds).
    ```
    {
      "key":"hello",
      "value":["there","yoii"],
      "ttl":500
    }
    ```
  * Response as:
      - Sucess : True/False
      - Message: Success message, Update message(if data was already present), Error message
      - Data: Respective data
    ```
    {
      "data": {
        "key": "hello",
        "value": [
            "there",
            "yoii"
        ],
        "ttl": 1718998846229906412
      },
      "message": "New value successfully added",
      "success": true
    }
    ```
  - **Get "/disconnect"**
  * Disconnects from the peer network and allows you to use it as a normal cache. The data will be shared only when you connect again to the server.
  
## Screenshots:
<br>
<center>
<img width="1000" src="https://github.com/Uttkarsh-raj/Gossip-Cache/assets/106571927/1534f7e3-afc4-47ca-9643-e3bef4d3dd79"></img>
<br>
<img width="1000" src="https://github.com/Uttkarsh-raj/Gossip-Cache/assets/106571927/01cd3706-568a-44e3-861c-653756395013"></img>
</center>
<br>

<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

<p align="right">(<a href="#readme-top">back to top</a>)</p>
