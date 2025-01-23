package build

const LocalBuildContent = `version: '3'

services:
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
