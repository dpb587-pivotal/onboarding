---
title: Install VirtualBox
labels:
- core tools
type: story-feature
weight: 3
---

Install VirtualBox
### What?
[VirtualBox](https://www.virtualbox.org/manual/ch01.html) is a tool that creates and manages virtual machines (VMs) on your development environment. [PCF Dev](https://pivotal.io/pcf-dev) uses VirtualBox as its virtualizer, so you'll need to install VirtualBox in order to run PCF Dev.

### How?
**[Download the latest VirtualBox installer from virtualbox.org](https://www.virtualbox.org/wiki/Downloads)**

### Expected Result
When you run `VBoxManage --version`, your output looks something like `5.1.18r114002`

### Resources
[Blog post: containers vs. virtual machines, what’s the difference?](https://www.solidfire.com/blog/containers-vs-vms/)
[Internal Google Doc: PCF Linux Containers Overview](https://docs.google.com/a/pivotal.io/document/d/1QNcmQCrHZNAr17ULoYis_sDPNV0X4IUO1grFDBlQdO8/edit?usp=sharing)—by CF's Senior VP of Engineering Onsi Fakhouri. I really recommend this one, even if you don't get to it now.