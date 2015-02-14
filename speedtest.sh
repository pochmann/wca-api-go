for url in `cat speedtest.txt`; do
  echo
  echo $url
  ab -q -n 1000 $url | grep 'Time per request' | head -1
done
