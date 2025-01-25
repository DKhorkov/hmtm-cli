package dockercompose

const ProdContent = `version: '3'

services:
  hmtm_<service-name>:
    container_name: hmtm_<service-name>
    image: hmtm_<service-name>
    build:
      context: ../../..
      dockerfile: ./build/package/Dockerfile
    ports:
      - "${HMTM_<service-name-upper>_OUTER_PORT}:${HMTM_<service-name-upper>_INNER_PORT}"
    depends_on:
      - hmtm_<service-name>_database
    volumes:
      - ../../../logs/:/app/logs/
    networks:
      - hmtm_network

  hmtm_<service-name>_database:
    container_name: hmtm_<service-name>_database
    hostname: hmtm_<service-name>_database
    image: postgres
    restart: always
    env_file:
      - ../../../.env
    volumes:
      - ../../../postgres_data:/var/lib/postgresql/data
      - ../../../postgres_backups:/backups
      - ../../../scripts/postgres:/scripts
    ports:
      - "${HMTM_<service-name-upper>_DB_OUTER_PORT}:${HMTM_<service-name-upper>_DB_INNER_PORT}"
    networks:
      - hmtm_network

networks:
  hmtm_network:
`
