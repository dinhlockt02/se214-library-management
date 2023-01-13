FROM migrate/migrate
COPY ./db/migration /migrations
ENTRYPOINT ["migrate"]
CMD ["--help"]