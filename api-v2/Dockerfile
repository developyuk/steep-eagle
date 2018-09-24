FROM python:3

WORKDIR /app

# RUN apt-get update && apt-get install -y --no-install-recommends \
# 		cron \
# 	&& rm -rf /var/lib/apt/lists/*
# COPY crontab /etc/cron.d/steepeagle
# RUN chmod 0600 /etc/cron.d/steepeagle

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt


CMD [ "python", "./main.py" ]

