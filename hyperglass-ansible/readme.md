# Hyperglass Ansible Deployment (Docker)

Развёртывание сервиса hyperglass (looking glass) на Rocky Linux 9
с использованием Ansible и Docker.

---

## Требования

- Ansible ≥ 2.12
- SSH-доступ к серверу
- Root-пароль
- Чистая Rocky Linux 9
- Место на диске 15гб.

---

## Запуск

1. Указать IP сервера в inventory.ini

2. Запустить:
```bash
ansible-playbook -i inventory.ini playbook.yml -u root --ask-pass

3. После прогона подождать 3-4 мин.