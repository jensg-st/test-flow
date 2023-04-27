FROM python

RUN pip install requests

COPY function.py /
RUN chmod 755 function.py

CMD ["python", "/function.py"]