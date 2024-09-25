# #!/bin/bash

# # Check if the .env file exists
# if [ ! -f ".env" ]; then
#     echo ".env file not found!"
#     exit 1
# fi

# # Read the .env file line by line
# while IFS= read -r line; do
#     # Ignore empty lines and comments
#     [[ -z "$line" || "$line" =~ ^# ]] && continue
    
#     # Export each line
#     export "$line"
# done < .env

# echo "Environment variables exported successfully."
#!/bin/bash

# Check if the .env file exists
if [ ! -f ".env" ]; then
    echo ".env file not found!"
    exit 1
fi

# Read the .env file line by line
while IFS= read -r line; do
    # Ignore empty lines and comments
    [[ -z "$line" || "$line" =~ ^# ]] && continue

    # Export each line
    export "$line"
done < .env

echo "Environment variables exported successfully."
