#!/bin/bash

# Script to downsize all WEBP images in current directory and subdirectories by 50%
# Usage: Save as downsize_webp_50.sh, chmod +x downsize_webp_50.sh, then ./downsize_webp_50.sh

find . -type f -name "*.webp" -print0 | while IFS= read -r -d '' file; do
    echo "Processing: $file"
    
    # Downsize by 50% (maintains aspect ratio), overwrites original
    convert "$file" -resize 50% "$file"
    
    echo "Downsized 50%: $file"
done

echo "50% downsizing complete for all WEBP files."
