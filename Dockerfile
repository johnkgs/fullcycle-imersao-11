FROM python:3.10.2-slim

RUN useradd -ms /bin/bash python

RUN pip install pipenv

USER python

WORKDIR /home/python/app

ENV PIPENV_VENV_IN_PROJECT=True


# Manter um container (fica lendo o dispositivo nulo do linux)
CMD ["tail", "-f", "/dev/null"] 
