# Temporal (Heavy WIP) [![GoDoc](https://godoc.org/github.com/RTradeLtd/Temporal/api?status.svg)](https://godoc.org/github.com/RTradeLtd/Temporal/api)

Temporal is an enterprise-grade storage solution featuring an easy to consume API that can be easily integrated into your existing application stack, providing all the benefits of the distributed web without any of the overhead that comes with running distributed storage nodes.  Initially targetting the public IPFS network, the next release cycle will bring support for additional protocols such as Ethereum Swarm and Private IPFS network. Temporal won't stop there and will continue to evolve as the space itself evolves. At RTrade we understand that the blockchain and distributed technology space evolve at an extremely quick pace so we are designing Temporal with modularity in mind so that as the space evolves we can evolve seamlessly with it.

Temporal's API comes in two flavours, hosted or on-site. Should you not have the resources to run your own distributed storage nodes and infrastructure, you can take advantage of our hosted API allowing us to manage all the storage nodes and infrastructure. All that you have to worry about is your application and using our API while we deal with all the hardware and infrastructure, allowing you to spend your time focusing on releasing products, rather than troubleshooting infrastructure failures which drain critical development resources. If you have the infrastructure and technical resources, you can also take advantage of Temporal being open source and deploy your own Temporal infrastructure. For on-site deployments we offer special paid for installations, maintenance, upgrades, and product usage information sessions allowing you to take full advantage of the power that comes with running your own Temporal infrastructure.


# Supported Technologies

The following is a list of distributed and decentralized storage technologies that Temporal intends to support.

IPFS (50% complete):

    Temporal supports integration with the public IPFS network and will evolve to support new features added to IPFS so that you will have the most optimal experience possible with an API that is constantly in development. 

    Soon after release, support for Private IPFS networks will be integrated into Temporal, allowing you to get the same benefits of the public IPFS network with the data security and privacy that comes with running a private network. This is extremely useful to financial institutions, data archivers, and other industries to whom data security and privacy is one of the primary concerns when integrating with any new technology.

Swarm (5% complete):

    Swarm blends the power of IPFS with the power of Blockchain and is common place in protocols like Ethereum. Temporal will provide an interface into the Ethereum mainnet Swarm protocol, allowing you to store your data onto the Ethereum blockchain. You'll also take advantage of having that particular data persistently stored on our high quality Ethereum node infrastructure.  Currently the Swarm protocol isn't fully integrated yet with Ethereum, and as such, the SWAP accounting protocol isn't fully implemented yet, which means nodes in the network aren't guaranteed to persist your data. Until SWAP is fully integrated we hope to offer increased utilization and adoption of SWARM by taking advantage of the data persistence offered by Temporal and our high quality Ethereum infrastructure.

STORJ (0% complete):

    Details TBA

SIA (0% complete):

    Details TBA

# Hosted API Early Access

We have an early access alpha of our hosted API for the public IPFS network, should you wish to test it please contact me at postables@rtradetechnologies.com

There is also a telegram room you can join https://t.me/RTradeTEMPORAL

# Temporal Administration

Temporal will be administered by a text-based interface, which will allow full control over all parts of a temporal instance, including spinning up additional infrastructure nodes to scale up as user demand increases.  The Text Based User interface is currently under construction, and is located in the `tui` folder. After the text based user interface is complete, and fully functional work on a graphical interface for Temporal administration will be done.

# Interacting with Temporal

There are two ways of interacting with Temporal (currently only the API is supported). One consisting of an API intended to be used by application developers, programmers, or organizations who want to integrate Temporal with their application stack. Second is a web client that can be used as a cloud storage platform for personal, and enterprise use cases (this will be done in a later release).

# System Monitoring

System monitoring is done using a combination of Zabbix, and Grafana. Zabbix is used for operating system, and hardware metrics (network utilization, resource utilization, storage utilization, etc...) while Grafana is used for API and application level information for things like IPFS node statistics, API requests, etc...

All templates, and grafana graphs, as well as necessary configuration filse and scripts to replicate our system monitoring backend are included in the `configs/` directory of this repository. The only configurations we don't include are our email configurations.

# Goals

* Provide an easy to use interface into distributed and decentralized storage technologies.
* Teach users how to use decentralized and distributed storage technologies.
* Introduce these new storage technologies to audiences who may have otherwise not heard of them.
* Help organizations make informed decisions about whether or not integrating distributed and decentralized storage technologies is the right thing to do for their business.

# Data Privacy

Privacy is a huge issue and concern for any form of cloud storage, but is rarely mentioned by projects promising revolutionary storage systems on a global scale. Many projects talk about all the technical benefits and features of their storage solutions which use IPFS, but data privacy isn't discussed. Data privacy is an extreme focus of RTrade, and we will not be releasing Temporal until we are absolutely certain that all data privacy laws we need to abide by are met.

# How're We Different

AT RTrade we made the concious decision not to launch an ICO and we will not be wasting our development efforts redesigning the wheel with any new storage protocol or blockchain solution. Although we're using bleeding edge technology, we're commited to using names and open source software that are already tested, and that have a thriving development community behind them. And finally, results matter; It is far to common in this space for companies to ask you to hand over your hand earned cash on the fleeting promise that it will lead to something, but that something is either never delivered, or extremely lack in features, and is not the original idea which was sold.

# Contributing Code

If you wish to contribute, create a new branch to do your work on, and submit a pull request to V2 branch.
The only requirement is that all code submitted must include tests.

# Repository Contents

* `api/`
    * This contains the API related code to temporal, as is the primary interface for interaction with temporal
    * all web frontends, applications, etc... use this api
* `api/middleware`
    * This is middleware used by the API and handles common functionality such as database connection parameters, rabbitmq parameters, etc...
* `bindings`
    * This is the go-ethereum bindings for the smart contracts that temporal uses
* `cli`
    * basic terminal-based cli application
* `configs`
    * parent directory contains systemd service files
* `configs/grafana`
    * JSON files for our grafana graphs
* `contracts`
    * solidity source code files for temporal smart contracts
* `database`
    * this is the package used by temporal for interacting with our database backend
* `docs`
    * contains all non-readme documentation
* `models`
    * models used by temporal
* `payments`
    * golang code for interacting with the payments smart contract
* `queue`
    * all queue related code for rabbitmq and ipfs pubsub queue
* `rtfs`
    * contains temporal code for interacting with ipfs nodes
* `rtfs_cluster`
    * contains temporal code for interacting with the ipfs cluster
* `server`
    * contains server related code for interacting with the ethereum blockchain
    * this will eventually be broken up into three seperate folders corresponding to each of the smart contracts
* `utils`
    * contains utility code used by temporal

# Funding

Currently the project is paid for out of pocket, and we will *not* be doing an ICO. Any funding received through donations, or private investment will be put to the following:

* Infrastructure
    * Servers
    * Hard drives
    * Misc infrastructure (networking, etc..)
* Code Audits
* Developer Tools
* Legal
    * Lawyers
    * Ensure we meet regulatory needs
* Media
    * Educational resources development
* Hiring additional talent for Temporal enterprise to bring project to completion

Should you wish to consider donations, or private investment email admin@rtradetechnologies.com

# Warnings

Until V1 is released, do not expect backwards compatability between non V1 versions, and V1 versions. If you are going to use temporal, please use vendoring to prevent unexpected breaks before V1

# Usage (to do)

## Usage: API

* `/api/v1/login`
    * This is used to login, and authenticate with temporal

* `/api/v1/register`
    * This is used to register a user account with temporal

* `/api/v1/ipfs/pin/:hash`
    * This is used to pin content to temporal
    * Note: this only pins to a local node

* `/api/v1/ipfs/add-file`
    * This is used to upload a file to IPFS through temporal
    * Note: this adds it first to a local node, followed by a cluster pin 

* `/api/v1/ipfs/remove-pin/:hash`
    * This is used to remove a pin from the local ipfs node
    * Note: only the admin of the temporal instance can call this route

* `/api/v1/ipfs-cluster/pin/:hash`
    * This is used to pin a hash to the cluster

* `/api/v1/ipfs-cluster/sync-errors-local`
    * This is used to parse through a local node cluster status, and sync all errors

* `/api/v1/ipfs-cluster/remove-pin/:hash`
    * This is used to remove a pin from the cluster
    * Note: this is onyl usable by the admin of a temporal instance

## Usage: Deploying Temporal
Authentication is done with JWT, for testing with postman see this link https://medium.com/@codebyjeff/using-postman-environment-variables-auth-tokens-ea9c4fe9d3d7


Setting up postman with the tests section:

    var data = JSON.parse(responseBody); // parses the response body
                                    // into json for us
    console.log(data);
    postman.setEnvironmentVariable("token", data.token);


1) install postgresql
2) setup database tables
3) run `scripts/temporal_install.sh`
4) Setup `config.json`

When running temporal the following services need to be ran:
    1) api
    2) queue-dpa
    3) queue-dfa
    4) ipfs-cluster-queue

Before running temporal for the first time, and after changing any of the models run `Temporal migrate`

# Tips and Tricks

Make sure you set the path to the ipfs repo, and ipfs-cluster directory in your `.bashrc` or other profile file.
