#!/usr/bin/env python3
"""
Parse .speakeasy/gen.lock and extract management.docVersion for go generate.
"""

import yaml
import sys
import os
import subprocess

def main():
    # Path to gen.lock file
    gen_lock_path = "swov1/.speakeasy/gen.lock"

    # Check if file exists
    if not os.path.exists(gen_lock_path):
        print(f"Error: {gen_lock_path} not found", file=sys.stderr)
        sys.exit(1)

    # Parse YAML file
    try:
        with open(gen_lock_path, 'r') as f:
            data = yaml.safe_load(f)
    except Exception as e:
        print(f"Error parsing {gen_lock_path}: {e}", file=sys.stderr)
        sys.exit(1)

    # Extract management.docVersion
    doc_version = data.get('management', {}).get('docVersion')

    if not doc_version:
        print("Error: management.docVersion not found in gen.lock", file=sys.stderr)
        sys.exit(1)

    print(f"Extracted docVersion: {doc_version}")

    # Read the template file
    template_path = "swov1/internal/hooks/version.go"
    try:
        with open(template_path, 'r') as f:
            template_content = f.read()
    except Exception as e:
        print(f"Error reading {template_path}: {e}", file=sys.stderr)
        sys.exit(1)

    # Replace the template variable with the actual docVersion
    updated_content = template_content.replace("{{.DocVersion}}", doc_version)

    # Write back the updated file
    try:
        with open(template_path, 'w') as f:
            f.write(updated_content)
        print(f"Updated {template_path} with docVersion: {doc_version}")
    except Exception as e:
        print(f"Error writing {template_path}: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()