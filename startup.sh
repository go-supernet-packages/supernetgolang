#!/bin/bash
GO=/usr/bin/go
for dir in $PWD/*
do
	case "${dir}" in 
    "$PWD/startup.sh"|"$PWD/supernetgolang")
        echo "matched"
    ;;
    *)
#    "${dir}/./startup.sh" stop 
    cd ${dir}
    # ${GO} env -w GOPRIVATE=gitlab.com/superbillpay/*
    ${GO} build *.go
    "${dir}/./startup.sh" $1
    cd ..
    ;;
    esac
done
