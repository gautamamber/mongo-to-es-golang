#!/bin/bash

# MongoDB connection details
MONGO_HOST="localhost"
MONGO_PORT="27017"
MONGO_DB="ECommerceDB"

# Collections and schema information (modify as needed)
collections=("users" "products" "orders" "customers" "inventory" "reviews" "transactions" "logs" "employees" "departments")

# Function to generate MongoDB shell commands to create collections and insert random data
generate_mongo_script() {
  echo "use $MONGO_DB;" > mongo_script.js
  for collection in "${collections[@]}"; do
    echo "db.createCollection('$collection');" >> mongo_script.js
    echo "for (let i = 0; i < 100; i++) {" >> mongo_script.js
    echo "  db.$collection.insertOne({" >> mongo_script.js
    echo "    name: 'Name' + i," >> mongo_script.js
    echo "    value: Math.floor(Math.random() * 1000)," >> mongo_script.js
    echo "    createdAt: new Date()" >> mongo_script.js
    echo "  });" >> mongo_script.js
    echo "}" >> mongo_script.js
  done
}

# Step 1: Generate MongoDB script with collection and data insertion
generate_mongo_script

# Step 2: Execute the script using the mongo shell
mongo --host $MONGO_HOST --port $MONGO_PORT $MONGO_DB < mongo_script.js

# Step 3: Cleanup the generated script
rm mongo_script.js

echo "Collections created, and data inserted successfully!"
