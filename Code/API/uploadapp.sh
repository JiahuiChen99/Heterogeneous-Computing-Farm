curl --location --request POST '/deploy?autodeploy=true&iot=true' \
--header 'Content-Type: multipart/form-data' \
--header 'Accept: application/json' \
--form 'app=@"/path/to/file"' \
--form 'sys_cpu_cores="12"' \
--form 'sys_gpu_cores="1000"' \
--form 'sys_ram="24576"' \
--form 'app_cpu_cores="8"' \
--form 'app_ram="8192"'