.PHONY: ;

# https://github.com/ansible/ansible/issues/42388
export ANSIBLE_CONFIG=ansible.cfg

INVENTORY_FILE = inventory.yaml

install:
	ansible-galaxy install -r requirements.yaml

playbook:
	ansible-playbook playbook.yaml -i $(INVENTORY_FILE)

ping:
	ansible all -m ping -i $(INVENTORY_FILE)
