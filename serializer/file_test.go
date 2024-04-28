package serializer

import (
	"github.com/stretchr/testify/assert"
	"pcbook/pb"
	"pcbook/sample"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	assert.Nil(t, err)
	//require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	assert.Nil(t, err)
	//require.NoError(t, err)
	assert.Equal(t, laptop1.Brand, laptop2.Brand)

	//require.True(t, proto.Equal(laptop1, laptop2))
}
