import os
import sys
import shutil

# Usage:
#   Before running tests or codegen locally, ensure you have set QUERY_START_TIME and QUERY_END_TIME
#   (e.g., by sourcing swov1/.env), then run:
#       python3 replace_env_in_yaml.py swov1/.speakeasy/tests.arazzo.yaml
# This replaces placeholders in the YAML with the current environment variable values.

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python replace_env_in_yaml.py <yaml_file>")
        sys.exit(1)
    yaml_file = sys.argv[1]
    backup_file = yaml_file + ".bak"
    shutil.copyfile(yaml_file, backup_file)

    start_time = os.environ.get("QUERY_START_TIME")
    end_time = os.environ.get("QUERY_END_TIME")
    if start_time is None or end_time is None:
        print("Environment variables QUERY_START_TIME and QUERY_END_TIME must be set", file=sys.stderr)
        sys.exit(1)

    with open(yaml_file, "r") as f:
        content = f.read()
    content = content.replace("__QUERY_START_TIME__", f'"{start_time}"')
    content = content.replace("__QUERY_END_TIME__", f'"{end_time}"')
    with open(yaml_file, "w") as f:
        f.write(content)
