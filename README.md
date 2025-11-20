# Как использовать

```powershell
Usage:
  perfomate [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  perfomance  Generate perfomance review
  self        Generate self review

Flags:
  -h, --help                 help for perfomate
  -i, --input-path string    input raw file path
      --json                 use JSON convertor
  -o, --output-path string   output path (application doesn't create folders) (default "./")
```

## Генерация perfomance-review

### Exel
```powershell
perfomate.exe perfomance -i .\perfomance.xlsx -o .\
```

### JSON
```powershell
perfomate.exe perfomance --json -i .\perfomance.json -o .\
```

## Генерация self-review

### Exel
```powershell
perfomate.exe self -i .\self.xlsx -o .\
```

### JSON
```powershell
perfomate.exe self --json -i .\self.json -o .\
```