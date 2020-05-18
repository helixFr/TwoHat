# Two Hat Security: freelancer contest

Contest entry by Helix [[profile]](https://www.freelancer.com/u/astaroht)

Basic implementation complete. Saw the contest just a few hours before its ending, still working on a few things.
## Performance Report

```bash
$ siege -c100 -t45S --content-type "application/json" '127.0.0.1:5050 GET {"text": "test 1"}'
```   


QPS/ Transaction Rate:  ~8500  
Memory usage:           ~535MB     
     


**For comparison (basic "Hello, world!" server)**  

```bash
$ siege -c100 -t45S '127.0.0.1:5050'
```

QPS/ Transaction Rate:  ~12500  
Memory usage:           ~15MB   

**System Specifications**   

2 cores, 4GB (virtual machine setup for optimal testing)  


 
## Installation Instructions

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install cffi.

```bash
pip install foobar
```  


**To run the server**  

Clone the repo, and in the directory run:

```bash
python3 server.py
```  
Server would start at localhost:5050 [[127.0.0.1:5050]](127.0.0.1:5050)

## How long it took?
**< 1 hour** :  Setup python server (flask) with Go backend (load_json, get_topics, etc.)  
**2 hours**  :    Attempting to get Go FastHttp to talk with Python (i.e. Go calls Python, then Python calls Go), boilerplate communication and runtime clash issue, scrapped.  
**< 1 hour** :  Setup Go net/http server (as in the given [article](https://blog.heroku.com/see_python_see_python_go_go_python_go))  

**Total** : ~ 4 hours  


## Miscellaneous


**Currently working** on two approaches for Go FastHttp (for contribution):
* One previously scrapped (Go calls Python, then Python calls Go, separate runtimes with boilerplate communication)
* Easy one: Python calls Go FastHttp, and also backend, using CFFI (i.e. initiation code is run by python)

## Scoring checkbox
* Conventional commits ✓
* Single puspose commits ✓
* README.md ✓
* Followed instructions ✓
* QPS position [idk]
* Code smell [please judge]
* Well documented code [please judge]
* Code broken into reasonable peices ✓
* 80% coverage ✓
* 100ms change dictionary ✗ (working)


## Credits

[shazow: gohttplib](https://github.com/shazow/gohttplib) For net/http wrapper from Go