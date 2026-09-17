echo -e "$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq -r --argjson id "$HERO_ID" ".[] | select(.id == \$id) | .connections.relatives")"
