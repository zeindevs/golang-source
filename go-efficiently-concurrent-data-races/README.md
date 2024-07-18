
- Optimistic concurrency control (OCC) is a concurrency management method that assumes conflicts between transaction are rate
- It allows transactions to proceed without locking or blocking.
- But before commiting changes at the DB layer, OCC checks for conflicts and rollbacks and discards if found.
- This approach enhances scalability and performance by reducing locking, enable applications to serve multiple users simultaneously.
