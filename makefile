cluster-create:
	k3d cluster create primary --agents 1

cluster-start:
	k3d cluster start primary

cluster-delete:
	k3d cluster delete primary