.PHONY: ;

INVENTORY_FILE = inventory.yaml

install:
	ansible-galaxy install -r requirements.yaml

playbook:
	ansible-playbook playbook.yaml -i $(INVENTORY_FILE)

ping:
	ansible all -m ping -i $(INVENTORY_FILE)
