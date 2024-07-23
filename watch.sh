which inotifywait > /dev/null || echo "inotifywait not found. Please install inotify-tools"

inotifywait -e close_write,moved_to,create -mr ./internal |
while read -r directory events filename; do
  echo "File $filename has been changed"
  cd gen
  go run .
  cd ..
  echo "Regen done"
  cd cmd
  go run .
  cd ..
  echo "Rebuild done"
done