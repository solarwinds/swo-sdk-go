import sys
import shutil
# Usage:
#   Before running tests or codegen locally, run:
#       python3 replace_env_in_yaml.py swov1/.speakeasy/tests.arazzo.yaml
# This replaces placeholders in the YAML with the current environment variable values.

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python replace_env_in_yaml.py <yaml_file>")
        sys.exit(1)
    yaml_file = sys.argv[1]
    backup_file = yaml_file + ".bak"
    shutil.copyfile(yaml_file, backup_file)

    now = datetime.now(timezone.utc) - timedelta(minutes=1)
    start_time = now - timedelta(hours=6)
    end_time = now

    start_time_str = start_time.strftime("%Y-%m-%dT%H:%M:%SZ")
    end_time_str = end_time.strftime("%Y-%m-%dT%H:%M:%SZ")

    with open(yaml_file, "r") as f:
        content = f.read()
    content = content.replace("__QUERY_START_TIME__", f'"{start_time_str}"')
    content = content.replace("__QUERY_END_TIME__", f'"{end_time_str}"')
    with open(yaml_file, "w") as f:
        f.write(content)
