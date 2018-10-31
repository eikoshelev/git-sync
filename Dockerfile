FROM alpine:3.8 

COPY  git-sync /bin/
CMD ["git-sync"] # in this case, environment variables will be used (if you want to reassign them, set the flags explicitly)
