FROM scratch

ARG SERVICE_NAME

WORKDIR /app

COPY bin/${SERVICE_NAME} /app/service

ENTRYPOINT [ "/app/service" ]