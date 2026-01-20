import sys
import shutil
from datetime import datetime, timedelta, timezone
import yaml

def replace_placeholders(obj, start_time_str, end_time_str):
    if isinstance(obj, dict):
        return {k: replace_placeholders(v, start_time_str, end_time_str) for k, v in obj.items()}
    elif isinstance(obj, list):
        return [replace_placeholders(i, start_time_str, end_time_str) for i in obj]
    elif isinstance(obj, str):
        if obj == "__QUERY_START_TIME__":
            return start_time_str
        elif obj == "__QUERY_END_TIME__":
            return end_time_str
        else:
            return obj
    else:
        return obj

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python replace_env_in_yaml.py <yaml_file>")
        sys.exit(1)
    yaml_file = sys.argv[1]
    backup_file = yaml_file + ".bak"
    shutil.copyfile(yaml_file, backup_file)

    now = datetime.now(timezone.utc) - timedelta(minutes=1)
    end_time = now

    start_time_str = "2026-01-12T21:00:00.026217086Z"
    end_time_str = end_time.strftime("%Y-%m-%dT%H:%M:%SZ")

    with open(yaml_file, "r") as f:
        data = yaml.safe_load(f)

    data = replace_placeholders(data, start_time_str, end_time_str)

    with open(yaml_file, "w") as f:
        yaml.safe_dump(data, f, default_flow_style=False)

    print(f"Successfully replaced placeholders in {yaml_file}")
