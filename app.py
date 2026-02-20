from random import randint
from flask import Flask
import logging

from opentelemetry import trace
from opentelemetry import metrics

tracer = trace.get_tracer("diceroller.tracer")
meter = metrics.get_meter("diceroller.metrics")

# 初始化计量器
roll_counter = meter.create_counter(
    "dice.rolls",
    description="The number of rolls by roll value"
)

app = Flask(__name__)
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


@app.route("/rolldice")
def roll_dice():
    return str(roll())


def roll():
    with tracer.start_as_current_span("roll") as roll_span:
        res = randint(1, 6)
        roll_span.set_attribute("res", res)
        roll_counter.add(1, {"res": res})
        return res


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)

