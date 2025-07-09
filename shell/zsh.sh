   # cmdr.ai Zsh integration
   export CMDRAI_ENABLED=1

   function precmd() {
     local last_status=$?
     export CMDRAI_LAST_EXIT=$last_status
     export CMDRAI_LAST_CMD=$(fc -ln -1)
     if [[ $last_status -ne 0 && "$CMDRAI_ENABLED" == "1" ]]; then
       cmdr.ai "$CMDRAI_LAST_CMD"
     fi
   }