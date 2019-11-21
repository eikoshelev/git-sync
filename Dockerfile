FROM alpine:3.8 

COPY  git-sync /bin/
CMD ["git-sync"]
