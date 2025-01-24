#!/bin/bash

deployment_dir=".deploy"

declare -A params=(
  [STACK]=""
  [IMAGE]=""
  [CADDY_HOST]=""
  [CADDY_TLS]=""
  [DEPLOY_DB_DSN_MYSQL_ELEVA_CONTABO]=""
  [DEPLOY_DB_DSN_MYSQL_ELEVA]=""
  [DEPLOY_JWT_SECRET]=""
  [DEPLOY_MAILER_EMAIL]=""
  [DEPLOY_MAILER_PASSWORD]=""
  [DEPLOY_MAILER_SMTP_HOST]=""
  [DEPLOY_MAILER_SMTP_PORT]=""
  ## add more envs here same on the debug.sh
)

for arg in "$@"; do
  case $arg in
  *=*)
    key="${arg%%=*}"
    value="${arg#*=}"

    if [[ -z "${params[$key]}" ]]; then
      params[$key]="$value"
      if [[ -z "$value" ]]; then
        echo "Missing value for $key"
        exit 1
      fi
    fi
    ;;
  esac
done
# Deploy the stack

yaml_content=$(cat "$deployment_dir/deployment.template.yml")

# exclude envs array strings
declare -A exclude_envs=(
  ["STACK"]=1
)

function replace_env() {
  if [[ -v ${exclude_envs[$1]} ]]; then
    return
  fi

  env_name="${1}_PLACEHOLDER"
  env_value="${2}"
  escaped_env_value=$(printf '%s' "$env_value" | sed 's/[&/\]/\\&/g')

  yaml_content=$(echo "$yaml_content" | sed "s|$env_name|$escaped_env_value|g")
}

# raplace all args on the template
for key in "${!params[@]}"; do
  replace_env "$key" "${params[$key]}"
done

stack="${params[STACK]}-${params[DEPLOYMENT_ENVIROMENT]}"
yaml_content=$(echo "$yaml_content" | sed "s|STACK_PLACEHOLDER|$stack|g")

echo "$yaml_content"
