FROM python

RUN pip install requests

COPY myfunc.py /
RUN chmod 755 myfunc.py

CMD ["python", "/myfunc.py"]