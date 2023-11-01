#! /bin/bash

# Function to display a progress bar
function progress_bar() {
  local duration=${1}


    already_done() { for ((done=0; done<$elapsed; done++)); do printf "▇"; done }
    remaining() { for ((remain=$elapsed; remain<$duration; remain++)); do printf " "; done }
    percentage() { printf "| %s%%" $(( (($elapsed)*100)/($duration)*100/100 )); }
    clean_line() { printf "\r"; }

  for (( elapsed=1; elapsed<=$duration; elapsed++ )); do
      already_done; remaining; percentage
      sleep 1
      clean_line
  done
  clean_line
}

function cassandra_install() {
    echo "Installing cassandra"
    if [ -x "$(command -v docker)" ]; then
        docker run --name some-cassandra -p 9042:9042 -d cassandra:latest
    else
        echo "Docker is not installed on your system"
        exit 1
    fi

    # Check if Cassandra is up and running
    echo "Waiting for Cassandra to start..."
    while true; do
        docker ps | grep some-cassandra | grep "Up"

        if [ $? -eq 0 ]; then
            
            break
        fi

        echo "Cassandra is not yet running. Retrying in 5 seconds..."
        sleep 5
    done
    progress_bar 70
    if [ -x "$(command -v snapd)" ]; then
        echo "Installing cqlsh"
        sudo snap install cqlsh
    else
        echo "Installng using pip"
        pip install cqlsh
    fi
}
    

function cassandra_create_keyspace(){
    docker exec -it some-cassandra cqlsh -e "CREATE KEYSPACE IF NOT EXISTS bookstore WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"
}


function cassandra_books_table() {
    docker exec -it some-cassandra cqlsh -e "create table bookstore.books(id UUID, title text, author text, pages int, publisher text, PRIMARY KEY(title));"
}


function main() {
    cassandra_install
    cassandra_create_keyspace
    cassandra_books_table
}


main