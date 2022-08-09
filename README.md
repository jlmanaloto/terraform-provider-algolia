# Terraform Provider Algolia

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

This repository is a [Terraform](https://www.terraform.io) Provider for Algolia forked from [philippe-vandermoere's](https://github.com/philippe-vandermoere) [terraform-provider-algolia](https://github.com/philippe-vandermoere/terraform-provider-algolia).

The code structure of this Terraform provider for Algolia has been modified to work with Pulumi's [Terraform Bridge Provider Boilerplate](https://github.com/pulumi/pulumi-tf-provider-boilerplate).

See [usage](https://registry.terraform.io/providers/philippe-vandermoere/algolia/latest/docs).

## Developer

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.16

### Installation

Clone repository.

```bash
$ git clone https://github.com/jlmanaloto/terraform-provider-algolia
```
To compile, change directory to ``terraform-provider-algolia`` and run ``make install``.

```bash
$ cd terraform-provider-algolia && make install
```
