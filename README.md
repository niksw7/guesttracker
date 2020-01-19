# guesttracker


To inject linkerd proxies
kubectl get -n hackerspace deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
