#!/bin/bash

trap "exit 1" TERM
TOP_PID=$$

if [ "$1" = "--help" ]; then
  echo "-d describe about this script"
  kill -s TERM $TOP_PID
fi 

if [ "$1" = "-d" ]; then
  echo "#####################################################"
  echo "## this script will work to building go-app,       ##"
  echo "## directory after builded is build and within,    ##"
  echo "## build directory has [appName].tar.gz            ##"
  echo "#####################################################"
  kill -s TERM $TOP_PID
fi 

IFS='/' # space is set as delimiter

read -ra ADDR <<< "$(pwd)" # str is read into an array as tokens separated by IFS

isKeep="true"

# find position of gopath
for i in ${ADDR[@]} ;
do

  if [ "$i" != "" ]; then 
    if [ "$isKeep" = "true" ]; then
      ((state++))
    fi
  fi

  if [ "$i" = "go" ]; then
    inGo=$i
  fi

  if [ "$i" = "src" ]; then
    inSrc=$i
  fi
  
  result="${inGo}${inSrc}"

  if [ "$result" = "gosrc" ]; then
    isKeep="false"
      
  fi


done


for  ((i=1; i<${state}; i++ )) ;
do 
      gopath="$gopath/${ADDR[$i]}"
done

export GOPATH="$gopath"
export AMCO_HOME=$(pwd)



# build
rm -rf build
mkdir ./build
go build -tags prod -o build/${ADDR[${#ADDR[@]}-1]} . && \
  cp -R ./serverConfig ./build && \
  cd build && \
  tar -czf ${ADDR[${#ADDR[@]}-1]}.tar.gz * && \
  rm -rf serverConfig && rm ${ADDR[${#ADDR[@]}-1]}







