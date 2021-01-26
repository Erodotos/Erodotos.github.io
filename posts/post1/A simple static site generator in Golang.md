# A simple static site generator in Golang

## Introduction

The purpose of this article is to show how I built an elementary **static site generator** using *Golang*. I know that several open-source static **site generators** are available on the internet; however, I have decided to create my own mini-project to learn how they work and satisfy my curiosity. Furthermore, this article can be a reference point to those who might try to create their own simplistic **static site generator**.  

First of all, let us clarify some basic terminology. The term static web page refers to a web page delivered to the users' browser, as stored on a server (e.g., plain .html files). On the other hand, a  dynamic web page is generated on demand by a web application when users seek to access that page.

In the next section, I will explain the principal concepts of a **static site generator** and its major tasks. Subsequently, we will see implementation details and dive into some *Golang* code. 

## Static Site Generator : How it works

The structure of a static site generator is straightforward. Firstly, we need to prepare a template for our website and provide some data to fill in. We forward these sources into the main engine, whereas a program produces the final website according to the given parameters. We can directly host this website on a server and put our content online.

![Static Site Generator : How it works](https://cloud.netlifyusercontent.com/assets/344dbf88-fdf9-42bb-adb4-46f01eedd629/da1ef4c9-9d18-49c4-9d01-2defed1af3df/ssg-ssr-01-ssg.png)


## Implementation

## Conclusions

## References