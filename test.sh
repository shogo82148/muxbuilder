ROOT=$(cd $(dirname $0); pwd)
DATA=`ls -1 $ROOT/tests`
for i in ${DATA}
do
    cd $ROOT/tests/$i
    go run setup.go
    go test
done
