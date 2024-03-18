FROM ubuntu:latest
LABEL authors="vladislavtrofimov"

ENTRYPOINT ["top", "-b"]