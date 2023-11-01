# TF Mod

A simple CLI utlity to create a folder structure for a new Terraform module.

from your `modules` folder run:

```bash
tfmod <module-name>
```

This will create a folder structure like this:

```bash
<module-name>
├── README.md
├── main.tf
├── outputs.tf
├── variables.tf
└── providers.tf
``` 