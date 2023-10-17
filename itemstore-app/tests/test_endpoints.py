import requests


def get_all_items():
    url = "http://localhost:8080/items"
    response = requests.get(url)

    if response.status_code == 200:
        items = response.json()
        for item in items:
            print(f"ID: {item['id']}, Name: {item['name']}, Price: {item['price']}")
    else:
        print(f"Failed to retrieve items. Status code: {response.status_code}")


def get_specific_item_id(id):
    item_id = str(id)  # Replace with the desired item ID
    url = f"http://localhost:8080/items/{item_id}"
    response = requests.get(url)

    if response.status_code == 200:
        item = response.json()
        print(f"ID: {item['id']}, Name: {item['name']}, Price: {item['price']}")
    elif response.status_code == 404:
        print(f"Item with ID {item_id} not found.")
    else:
        print(f"Failed to retrieve the item. Status code: {response.status_code}")


def create_item(id):
    new_item = {
        "id": str(id),  # Replace with a unique ID
        "name": "New Item",
        "price": 9.99,
    }

    url = "http://localhost:8080/items"
    response = requests.post(url, json=new_item)

    if response.status_code == 200:
        created_item = response.json()
        print(f"Item created with ID: {created_item['id']}")
    else:
        print(f"Failed to create the item. Status code: {response.status_code}")


def update_item(id):
    updated_item = {
        "id": str(id),  # Replace with the ID of the item you want to update
        "name": "Updated Item",
        "price": 19.99,
    }

    url = f'http://localhost:8080/items/{updated_item["id"]}'
    response = requests.put(url, json=updated_item)

    if response.status_code == 200:
        updated_item = response.json()
        print(f"Item updated with ID: {updated_item['id']}")
    elif response.status_code == 404:
        print(f"Item with ID {updated_item['id']} not found.")
    else:
        print(f"Failed to update the item. Status code: {response.status_code}")


def delete_item(id):
    item_id = str(id)  # Replace with the ID of the item you want to delete
    url = f"http://localhost:8080/items/{item_id}"
    response = requests.delete(url)

    if response.status_code == 204:
        print(f"Item with ID {item_id} has been deleted.")
    elif response.status_code == 404:
        print(f"Item with ID {item_id} not found.")
    else:
        print(f"Failed to delete the item. Status code: {response.status_code}")


if __name__ == "__main__":
    create_item(1)
    create_item(2)
    create_item(3)
    get_specific_item_id(3)
    update_item(2)
    update_item(1)
    get_all_items()
    delete_item(3)
