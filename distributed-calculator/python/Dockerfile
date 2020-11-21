FROM python:3.7.1-alpine3.8
COPY . /app
WORKDIR /app
RUN pip install flask flask_cors
RUN pip install newrelic
ADD newrelic.ini .
ENV NEW_RELIC_CONFIG_FILE=newrelic.ini
ENV NEW_RELIC_APP_NAME="distributed-calculator-python-multiplier"
ENTRYPOINT ["newrelic-admin", "run-program"]
EXPOSE 5000
CMD ["python", "app.py"]
