import requests
import random
import time

BASE_URL = "http://localhost:8000"

try:
    while True:

        # Generate product list traffic
        response = requests.get(
            f"{BASE_URL}/products"
        )

        print(f"GET /products -> {response.status_code}")

        # Generate occasional failures
        if random.randint(1, 5) == 1:

            response = requests.get(
                f"{BASE_URL}/products?fail=true"
            )

            print(
                f"GET /products?fail=true -> {response.status_code}"
            )

        # Generate product creation traffic
        response = requests.post(
            f"{BASE_URL}/products",
            json={
                "id": f"p{random.randint(1000, 9999)}",
                "name": f"Product-{random.randint(1, 100)}",
                "price": round(
                    random.uniform(10, 100),
                    2,
                ),
            },
        )

        print(
            f"POST /products -> {response.status_code}"
        )

        time.sleep(2)

except KeyboardInterrupt:
    print("Traffic generator stopped")