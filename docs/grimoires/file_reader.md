# FileReader Grimoire

The `FileReader` grimoire provides a unified interface for reading and parsing various file formats in Carrion. It automatically detects file types based on extension and returns appropriately structured data.

## Supported File Types

| Type | Extensions | Return Type | Description |
|------|------------|-------------|-------------|
| `txt` | `.txt`, `.text`, `.log`, `.md`, `.rst` | Array of strings | Each line as an element |
| `csv` | `.csv`, `.tsv` | Array of arrays | Rows of cells |
| `json` | `.json` | Hash or Array | Parsed JSON structure |
| `yaml` | `.yaml`, `.yml` | Hash or Array | Parsed YAML structure |
| `toml` | `.toml` | Hash | Parsed TOML structure |
| `xml` | `.xml` | Hash | Nested structure with attributes |
| `ini` | `.ini`, `.cfg`, `.conf` | Hash | Sections containing key-value pairs |
| `properties` | `.properties` | Hash | Java-style key-value pairs |
| `excel` | `.xlsx`, `.xls`, `.xlsm`, `.xlsb` | Array of arrays | 2D array of cell values |

## Basic Usage

### Auto-Detection

FileReader automatically detects the file type from the extension:

```carrion
# JSON file
reader = FileReader("config.json")
config = reader.read()
print(config["database"]["host"])

# YAML file
reader = FileReader("docker-compose.yml")
services = reader.read()

# CSV file
reader = FileReader("users.csv")
rows = reader.read()
for row in rows:
    print(row)
```

### Explicit Type

You can explicitly specify the file type:

```carrion
# Treat a .txt file as CSV
reader = FileReader("data.txt", "csv")
rows = reader.read()

# Force JSON parsing
reader = FileReader("api_response", "json")
data = reader.read()
```

## Initialization

```carrion
reader = FileReader(filename, file_type="")
```

**Parameters:**
- `filename` (str): Path to the file to read
- `file_type` (str, optional): Explicit file type. If empty, auto-detected from extension.

**Example:**
```carrion
# Auto-detect
reader = FileReader("data.csv")

# Explicit type
reader = FileReader("data.txt", "csv")
```

## Methods

### Universal Methods

#### `read()`

Reads the file using the appropriate parser for its type.

```carrion
reader = FileReader("config.yaml")
data = reader.read()
```

**Returns:** Data structure appropriate for the file type.

---

#### `read_with_headers()`

Reads CSV or Excel files and returns an array of maps using the first row as headers.

```carrion
reader = FileReader("employees.csv")
employees = reader.read_with_headers()

for emp in employees:
    print(f"{emp['name']} - {emp['department']}")
```

**Returns:** Array of Hash objects, or `None` if not applicable.

---

#### `close()`

Closes any open file handles (primarily for Excel files).

```carrion
reader = FileReader("data.xlsx")
data = reader.read()
reader.close()
```

---

#### `get_type()`

Returns the detected or specified file type.

```carrion
reader = FileReader("config.yaml")
print(reader.get_type())  # Output: yaml
```

---

#### `get_supported_types()`

Returns an array of all supported file type names.

```carrion
reader = FileReader("any.txt")
types = reader.get_supported_types()
# ["txt", "csv", "json", "yaml", "toml", "xml", "ini", "properties", "excel"]
```

---

#### `get_extensions_for_type(type_name)`

Returns the file extensions associated with a type.

```carrion
reader = FileReader("any.txt")
exts = reader.get_extensions_for_type("yaml")
# [".yaml", ".yml"]
```

---

#### `is_valid_type()`

Checks if the file extension matches the declared type.

```carrion
reader = FileReader("data.json", "json")
print(reader.is_valid_type())  # True

reader = FileReader("data.txt", "json")
print(reader.is_valid_type())  # False
```

### Configuration Methods

#### `set_delimiter(delimiter)`

Sets the delimiter for CSV parsing. Returns self for method chaining.

```carrion
# Pipe-delimited file
reader = FileReader("data.csv")
reader.set_delimiter("|")
rows = reader.read()

# Method chaining
rows = FileReader("data.csv").set_delimiter(";").read()
```

**Note:** `.tsv` files automatically use tab (`\t`) as delimiter.

---

#### `set_encoding(encoding)`

Sets the character encoding for file reading. Returns self for method chaining.

```carrion
reader = FileReader("legacy.csv")
reader.set_encoding("latin-1")
data = reader.read()

# Method chaining
data = FileReader("data.csv").set_encoding("utf-16").read()
```

**Default:** `utf-8`

### Text File Methods

#### `read_txt()` / `read_lines()`

Reads a text file and returns all lines as an array.

```carrion
reader = FileReader("log.txt")
lines = reader.read_txt()

for line in lines:
    print(line)
```

---

#### `process_lines(callback)`

Processes each line with a callback function. Memory efficient for large files.

```carrion
reader = FileReader("large_file.txt")

count = reader.process_lines(spell(line):
    if line.contains("ERROR"):
        print(line)
)

print(f"Processed {count} lines")
```

**Returns:** Number of lines processed.

### CSV File Methods

#### `read_csv()`

Reads a CSV file and returns rows as arrays.

```carrion
reader = FileReader("data.csv")
rows = reader.read_csv()

# First row is headers
headers = rows[0]
for i in range(1, len(rows)):
    print(rows[i])
```

---

#### `read_csv_with_headers()`

Reads CSV and returns array of maps using first row as headers.

```carrion
reader = FileReader("users.csv")
users = reader.read_csv_with_headers()

for user in users:
    print(f"Name: {user['name']}, Email: {user['email']}")
```

### JSON File Methods

#### `read_json()`

Reads and parses a JSON file.

```carrion
reader = FileReader("config.json")
config = reader.read_json()

print(config["server"]["port"])
print(config["features"][0])
```

### YAML File Methods

#### `read_yaml()`

Reads and parses a YAML file.

```carrion
reader = FileReader("docker-compose.yml")
compose = reader.read_yaml()

for service_name, service_config in pairs(compose["services"]):
    print(f"Service: {service_name}")
    print(f"  Image: {service_config['image']}")
```

### TOML File Methods

#### `read_toml()`

Reads and parses a TOML file.

```carrion
reader = FileReader("Cargo.toml")
cargo = reader.read_toml()

print(f"Package: {cargo['package']['name']}")
print(f"Version: {cargo['package']['version']}")
```

### XML File Methods

#### `read_xml()`

Reads and parses an XML file into a nested Hash structure.

```carrion
reader = FileReader("config.xml")
xml = reader.read_xml()

# Access elements
print(xml["root"]["settings"]["#text"])

# Access attributes
print(xml["root"]["@attributes"]["version"])
```

**Structure:**
- Element names become keys
- `@attributes` contains element attributes
- `#text` contains text content
- Repeated elements become arrays

### INI File Methods

#### `read_ini()`

Reads and parses an INI/CFG file.

```carrion
reader = FileReader("config.ini")
config = reader.read_ini()

# Access by section
db_host = config["database"]["host"]
db_port = config["database"]["port"]

# Default section
app_name = config["DEFAULT"]["name"]
```

### Properties File Methods

#### `read_properties()`

Reads Java-style properties files.

```carrion
reader = FileReader("application.properties")
props = reader.read_properties()

print(props["server.port"])
print(props["database.url"])
```

### Excel File Methods

#### `read_excel(sheet_name="")`

Reads data from an Excel file. If no sheet name provided, reads the first sheet.

```carrion
reader = FileReader("report.xlsx")

# Read first sheet
data = reader.read_excel()

# Read specific sheet
sales = reader.read_excel("Sales")
```

---

#### `get_sheets()`

Returns list of all sheet names in the workbook.

```carrion
reader = FileReader("workbook.xlsx")
sheets = reader.get_sheets()

for sheet in sheets:
    print(f"Sheet: {sheet}")
```

---

#### `read_sheet(sheet_name)`

Reads all data from a specific sheet.

```carrion
reader = FileReader("data.xlsx")
sales_data = reader.read_sheet("Q1 Sales")
```

---

#### `read_row(sheet_name, row_num)`

Reads a specific row (1-based indexing).

```carrion
reader = FileReader("data.xlsx")
header_row = reader.read_row("Sheet1", 1)
first_data_row = reader.read_row("Sheet1", 2)
```

---

#### `read_cell(sheet_name, cell_ref)`

Reads a specific cell value.

```carrion
reader = FileReader("data.xlsx")
value = reader.read_cell("Sheet1", "A1")
total = reader.read_cell("Summary", "D15")
```

---

#### `read_excel_with_headers(sheet_name="")`

Reads Excel sheet and returns array of maps using first row as headers.

```carrion
reader = FileReader("employees.xlsx")
employees = reader.read_excel_with_headers("Staff")

for emp in employees:
    print(f"{emp['Name']} - {emp['Department']}")
```

---

#### `iterate_rows(callback, sheet_name="")`

Iterates over rows with a callback function.

```carrion
reader = FileReader("data.xlsx")

reader.iterate_rows(spell(row_num, row_data):
    print(f"Row {row_num}: {row_data}")
, "Sheet1")
```

## Complete Examples

### Processing a Configuration File

```carrion
# config.yaml
# database:
#   host: localhost
#   port: 5432
# features:
#   - auth
#   - logging

reader = FileReader("config.yaml")
config = reader.read()

db_host = config["database"]["host"]
db_port = config["database"]["port"]

print(f"Connecting to {db_host}:{db_port}")

for feature in config["features"]:
    print(f"Enabling feature: {feature}")
```

### CSV Data Processing

```carrion
# sales.csv
# date,product,quantity,price
# 2024-01-01,Widget,10,29.99
# 2024-01-02,Gadget,5,49.99

reader = FileReader("sales.csv")
sales = reader.read_with_headers()

total_revenue = 0
for sale in sales:
    qty = int(sale["quantity"])
    price = float(sale["price"])
    revenue = qty * price
    total_revenue = total_revenue + revenue
    print(f"{sale['date']}: {sale['product']} - ${revenue}")

print(f"Total Revenue: ${total_revenue}")
```

### Multi-Sheet Excel Report

```carrion
reader = FileReader("quarterly_report.xlsx")

# Get all sheets
sheets = reader.get_sheets()
print(f"Available sheets: {sheets}")

# Process each quarter
for sheet in sheets:
    if sheet.startswith("Q"):
        print(f"\n=== {sheet} ===")
        data = reader.read_excel_with_headers(sheet)

        for row in data:
            print(f"  {row['Region']}: ${row['Sales']}")

reader.close()
```

### Log File Analysis

```carrion
reader = FileReader("application.log")

error_count = 0
warning_count = 0

reader.process_lines(spell(line):
    if line.contains("ERROR"):
        error_count = error_count + 1
        print(f"ERROR: {line}")
    otherwise line.contains("WARN"):
        warning_count = warning_count + 1
)

print(f"\nSummary: {error_count} errors, {warning_count} warnings")
```

### INI Configuration with Sections

```carrion
# app.ini
# [server]
# host = 0.0.0.0
# port = 8080
#
# [database]
# driver = postgresql
# host = localhost

reader = FileReader("app.ini")
config = reader.read()

server = config["server"]
db = config["database"]

print(f"Server: {server['host']}:{server['port']}")
print(f"Database: {db['driver']}://{db['host']}")
```

## Error Handling

FileReader returns `None` when operations fail. Always check for None:

```carrion
reader = FileReader("maybe_missing.json")
data = reader.read()

if data == None:
    print("Failed to read file")
else:
    print(data)
```

## Best Practices

1. **Close Excel files** when done to release resources:
   ```carrion
   reader = FileReader("data.xlsx")
   data = reader.read()
   reader.close()
   ```

2. **Use `process_lines()` for large text files** to avoid loading everything into memory.

3. **Use `read_with_headers()`** for CSV/Excel files with header rows to get named access.

4. **Set encoding explicitly** for non-UTF-8 files:
   ```carrion
   reader = FileReader("legacy.csv").set_encoding("latin-1")
   ```

5. **Chain configuration methods** for cleaner code:
   ```carrion
   data = FileReader("data.csv").set_delimiter(";").set_encoding("utf-16").read()
   ```
