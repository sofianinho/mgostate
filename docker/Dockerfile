FROM scratch
MAINTAINER Sofiane Imadali <sofiane.imadali@orange.com>
LABEL Description="mgostate tests the existence of your mongodb server, database and collection. Use with --help for usage" Version="1.0"

# including the nobody user, inexistent in the scratch image
COPY etc /etc
# including the mgostate program, statically linked (no lib dependencies)
ADD mgostate /app/mgostate
# no need to be a root for this tool
USER nobody
# read the help message for the commands and flags
ENTRYPOINT ["/app/mgostate"]
CMD ["--help"]
