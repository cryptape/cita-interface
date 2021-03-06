// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package blockchain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProofType int32

const (
	ProofType_AuthorityRound ProofType = 0
	ProofType_Raft           ProofType = 1
	ProofType_Tendermint     ProofType = 2
)

var ProofType_name = map[int32]string{
	0: "AuthorityRound",
	1: "Raft",
	2: "Tendermint",
}
var ProofType_value = map[string]int32{
	"AuthorityRound": 0,
	"Raft":           1,
	"Tendermint":     2,
}

func (x ProofType) String() string {
	return proto.EnumName(ProofType_name, int32(x))
}
func (ProofType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{0}
}

type Crypto int32

const (
	Crypto_SECP Crypto = 0
	Crypto_SM2  Crypto = 1
)

var Crypto_name = map[int32]string{
	0: "SECP",
	1: "SM2",
}
var Crypto_value = map[string]int32{
	"SECP": 0,
	"SM2":  1,
}

func (x Crypto) String() string {
	return proto.EnumName(Crypto_name, int32(x))
}
func (Crypto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{1}
}

type Proof struct {
	Content              []byte    `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Type                 ProofType `protobuf:"varint,2,opt,name=type,enum=ProofType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (dst *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(dst, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Proof) GetType() ProofType {
	if m != nil {
		return m.Type
	}
	return ProofType_AuthorityRound
}

type BlockHeader struct {
	Prevhash             []byte   `protobuf:"bytes,1,opt,name=prevhash,proto3" json:"prevhash,omitempty"`
	Timestamp            uint64   `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	StateRoot            []byte   `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TransactionsRoot     []byte   `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	ReceiptsRoot         []byte   `protobuf:"bytes,6,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`
	GasUsed              uint64   `protobuf:"varint,7,opt,name=gas_used,json=gasUsed" json:"gas_used,omitempty"`
	GasLimit             uint64   `protobuf:"varint,8,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
	Proof                *Proof   `protobuf:"bytes,9,opt,name=proof" json:"proof,omitempty"`
	Proposer             []byte   `protobuf:"bytes,10,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevhash() []byte {
	if m != nil {
		return m.Prevhash
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTransactionsRoot() []byte {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *BlockHeader) GetReceiptsRoot() []byte {
	if m != nil {
		return m.ReceiptsRoot
	}
	return nil
}

func (m *BlockHeader) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *BlockHeader) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *BlockHeader) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *BlockHeader) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type Status struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{2}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Status) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type AccountGasLimit struct {
	CommonGasLimit       uint64            `protobuf:"varint,1,opt,name=common_gas_limit,json=commonGasLimit" json:"common_gas_limit,omitempty"`
	SpecificGasLimit     map[string]uint64 `protobuf:"bytes,2,rep,name=specific_gas_limit,json=specificGasLimit" json:"specific_gas_limit,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountGasLimit) Reset()         { *m = AccountGasLimit{} }
func (m *AccountGasLimit) String() string { return proto.CompactTextString(m) }
func (*AccountGasLimit) ProtoMessage()    {}
func (*AccountGasLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{3}
}
func (m *AccountGasLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGasLimit.Unmarshal(m, b)
}
func (m *AccountGasLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGasLimit.Marshal(b, m, deterministic)
}
func (dst *AccountGasLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGasLimit.Merge(dst, src)
}
func (m *AccountGasLimit) XXX_Size() int {
	return xxx_messageInfo_AccountGasLimit.Size(m)
}
func (m *AccountGasLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGasLimit.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGasLimit proto.InternalMessageInfo

func (m *AccountGasLimit) GetCommonGasLimit() uint64 {
	if m != nil {
		return m.CommonGasLimit
	}
	return 0
}

func (m *AccountGasLimit) GetSpecificGasLimit() map[string]uint64 {
	if m != nil {
		return m.SpecificGasLimit
	}
	return nil
}

type RichStatus struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Nodes                [][]byte `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Interval             uint64   `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RichStatus) Reset()         { *m = RichStatus{} }
func (m *RichStatus) String() string { return proto.CompactTextString(m) }
func (*RichStatus) ProtoMessage()    {}
func (*RichStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{4}
}
func (m *RichStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RichStatus.Unmarshal(m, b)
}
func (m *RichStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RichStatus.Marshal(b, m, deterministic)
}
func (dst *RichStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RichStatus.Merge(dst, src)
}
func (m *RichStatus) XXX_Size() int {
	return xxx_messageInfo_RichStatus.Size(m)
}
func (m *RichStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RichStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RichStatus proto.InternalMessageInfo

func (m *RichStatus) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *RichStatus) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RichStatus) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *RichStatus) GetInterval() uint64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type Transaction struct {
	To                   string   `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	Quota                uint64   `protobuf:"varint,3,opt,name=quota" json:"quota,omitempty"`
	ValidUntilBlock      uint64   `protobuf:"varint,4,opt,name=valid_until_block,json=validUntilBlock" json:"valid_until_block,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Value                uint64   `protobuf:"varint,6,opt,name=value" json:"value,omitempty"`
	ChainId              uint32   `protobuf:"varint,7,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	Version              uint32   `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{5}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *Transaction) GetQuota() uint64 {
	if m != nil {
		return m.Quota
	}
	return 0
}

func (m *Transaction) GetValidUntilBlock() uint64 {
	if m != nil {
		return m.ValidUntilBlock
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type UnverifiedTransaction struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Signature            []byte       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Crypto               Crypto       `protobuf:"varint,3,opt,name=crypto,enum=Crypto" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnverifiedTransaction) Reset()         { *m = UnverifiedTransaction{} }
func (m *UnverifiedTransaction) String() string { return proto.CompactTextString(m) }
func (*UnverifiedTransaction) ProtoMessage()    {}
func (*UnverifiedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{6}
}
func (m *UnverifiedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnverifiedTransaction.Unmarshal(m, b)
}
func (m *UnverifiedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnverifiedTransaction.Marshal(b, m, deterministic)
}
func (dst *UnverifiedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnverifiedTransaction.Merge(dst, src)
}
func (m *UnverifiedTransaction) XXX_Size() int {
	return xxx_messageInfo_UnverifiedTransaction.Size(m)
}
func (m *UnverifiedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_UnverifiedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_UnverifiedTransaction proto.InternalMessageInfo

func (m *UnverifiedTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *UnverifiedTransaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *UnverifiedTransaction) GetCrypto() Crypto {
	if m != nil {
		return m.Crypto
	}
	return Crypto_SECP
}

type SignedTransaction struct {
	TransactionWithSig   *UnverifiedTransaction `protobuf:"bytes,1,opt,name=transaction_with_sig,json=transactionWithSig" json:"transaction_with_sig,omitempty"`
	TxHash               []byte                 `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Signer               []byte                 `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignedTransaction) Reset()         { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()    {}
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{7}
}
func (m *SignedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransaction.Unmarshal(m, b)
}
func (m *SignedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransaction.Marshal(b, m, deterministic)
}
func (dst *SignedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransaction.Merge(dst, src)
}
func (m *SignedTransaction) XXX_Size() int {
	return xxx_messageInfo_SignedTransaction.Size(m)
}
func (m *SignedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransaction proto.InternalMessageInfo

func (m *SignedTransaction) GetTransactionWithSig() *UnverifiedTransaction {
	if m != nil {
		return m.TransactionWithSig
	}
	return nil
}

func (m *SignedTransaction) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *SignedTransaction) GetSigner() []byte {
	if m != nil {
		return m.Signer
	}
	return nil
}

type BlockBody struct {
	Transactions         []*SignedTransaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{8}
}
func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (dst *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(dst, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetTransactions() []*SignedTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Block struct {
	Version              uint32       `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Header               *BlockHeader `protobuf:"bytes,2,opt,name=header" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{9}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlockWithProof struct {
	Blk                  *Block   `protobuf:"bytes,1,opt,name=blk" json:"blk,omitempty"`
	Proof                *Proof   `protobuf:"bytes,2,opt,name=proof" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockWithProof) Reset()         { *m = BlockWithProof{} }
func (m *BlockWithProof) String() string { return proto.CompactTextString(m) }
func (*BlockWithProof) ProtoMessage()    {}
func (*BlockWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{10}
}
func (m *BlockWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockWithProof.Unmarshal(m, b)
}
func (m *BlockWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockWithProof.Marshal(b, m, deterministic)
}
func (dst *BlockWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockWithProof.Merge(dst, src)
}
func (m *BlockWithProof) XXX_Size() int {
	return xxx_messageInfo_BlockWithProof.Size(m)
}
func (m *BlockWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_BlockWithProof proto.InternalMessageInfo

func (m *BlockWithProof) GetBlk() *Block {
	if m != nil {
		return m.Blk
	}
	return nil
}

func (m *BlockWithProof) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type BlockTxs struct {
	Height               uint64     `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Body                 *BlockBody `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockTxs) Reset()         { *m = BlockTxs{} }
func (m *BlockTxs) String() string { return proto.CompactTextString(m) }
func (*BlockTxs) ProtoMessage()    {}
func (*BlockTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_95db5865cb60afa1, []int{11}
}
func (m *BlockTxs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTxs.Unmarshal(m, b)
}
func (m *BlockTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTxs.Marshal(b, m, deterministic)
}
func (dst *BlockTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTxs.Merge(dst, src)
}
func (m *BlockTxs) XXX_Size() int {
	return xxx_messageInfo_BlockTxs.Size(m)
}
func (m *BlockTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTxs.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTxs proto.InternalMessageInfo

func (m *BlockTxs) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockTxs) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Proof)(nil), "Proof")
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*AccountGasLimit)(nil), "AccountGasLimit")
	proto.RegisterMapType((map[string]uint64)(nil), "AccountGasLimit.SpecificGasLimitEntry")
	proto.RegisterType((*RichStatus)(nil), "RichStatus")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*UnverifiedTransaction)(nil), "UnverifiedTransaction")
	proto.RegisterType((*SignedTransaction)(nil), "SignedTransaction")
	proto.RegisterType((*BlockBody)(nil), "BlockBody")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*BlockWithProof)(nil), "BlockWithProof")
	proto.RegisterType((*BlockTxs)(nil), "BlockTxs")
	proto.RegisterEnum("ProofType", ProofType_name, ProofType_value)
	proto.RegisterEnum("Crypto", Crypto_name, Crypto_value)
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_95db5865cb60afa1) }

var fileDescriptor_blockchain_95db5865cb60afa1 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0x26, 0xf3, 0x93, 0x99, 0x9c, 0xe9, 0xce, 0xa6, 0xd6, 0xee, 0x12, 0x76, 0xf9, 0x59, 0x05,
	0x84, 0x56, 0x45, 0xca, 0x45, 0x41, 0x08, 0xb8, 0x6b, 0xab, 0x15, 0x45, 0x02, 0x69, 0xe5, 0x99,
	0x8a, 0xcb, 0x28, 0x4d, 0xdc, 0xc6, 0x74, 0x26, 0x0e, 0x89, 0xa7, 0x74, 0x5e, 0x80, 0x4b, 0x1e,
	0x8a, 0x67, 0x40, 0x3c, 0x0f, 0xc7, 0xc7, 0x49, 0xc7, 0x53, 0x2a, 0x21, 0xee, 0xf2, 0x7d, 0xfe,
	0x7c, 0xfc, 0xf9, 0xf8, 0x9c, 0x13, 0x08, 0x2f, 0x57, 0x2a, 0xbf, 0xc9, 0xcb, 0x4c, 0x56, 0x49,
	0xdd, 0x28, 0xad, 0xe2, 0x13, 0x18, 0xbf, 0x6b, 0x94, 0xba, 0x62, 0x11, 0x4c, 0x72, 0x55, 0x69,
	0x51, 0xe9, 0xc8, 0x7b, 0xed, 0xbd, 0x39, 0xe0, 0x3d, 0x64, 0x1f, 0xc3, 0x48, 0x6f, 0x6b, 0x11,
	0x0d, 0x90, 0x9e, 0x1f, 0x43, 0x42, 0xfa, 0x25, 0x32, 0x9c, 0xf8, 0xf8, 0xcf, 0x01, 0xcc, 0x4e,
	0x4d, 0xdc, 0x73, 0x91, 0x15, 0xa2, 0x61, 0x2f, 0x61, 0x5a, 0x37, 0xe2, 0xb6, 0xcc, 0xda, 0xb2,
	0x0b, 0x75, 0x8f, 0xd9, 0x87, 0x10, 0x68, 0xb9, 0x16, 0xad, 0xce, 0xd6, 0x35, 0x05, 0x1c, 0xf1,
	0x1d, 0xc1, 0x5e, 0x80, 0x5f, 0x0a, 0x79, 0x5d, 0xea, 0x68, 0x48, 0x4b, 0x1d, 0x62, 0x1f, 0x01,
	0xa0, 0x40, 0x8b, 0x14, 0x4f, 0xd6, 0xd1, 0x88, 0x62, 0x06, 0xc4, 0x70, 0x24, 0xd8, 0x17, 0x70,
	0xa8, 0x9b, 0xac, 0x6a, 0xb3, 0x5c, 0x4b, 0x55, 0xb5, 0x56, 0x35, 0x26, 0x55, 0xe8, 0x2e, 0x90,
	0xf8, 0x53, 0x78, 0xd2, 0x88, 0x5c, 0xc8, 0x5a, 0x77, 0x42, 0x9f, 0x84, 0x07, 0x3d, 0x49, 0xa2,
	0x0f, 0x60, 0x7a, 0x9d, 0xb5, 0xe9, 0xa6, 0x15, 0x45, 0x34, 0x21, 0x2b, 0x13, 0xc4, 0x17, 0x08,
	0xd9, 0x2b, 0x08, 0xcc, 0xd2, 0x4a, 0xae, 0xa5, 0x8e, 0xa6, 0xb4, 0x66, 0xb4, 0x3f, 0x1a, 0x8c,
	0xd7, 0x1b, 0xd7, 0x26, 0x3b, 0x51, 0x80, 0x0b, 0xb3, 0x63, 0xdf, 0xe6, 0x8a, 0x5b, 0xd2, 0x26,
	0x46, 0xd5, 0xaa, 0x15, 0x4d, 0x04, 0x7d, 0x62, 0x2c, 0x8e, 0xbf, 0x02, 0x7f, 0x81, 0x17, 0xda,
	0xb4, 0x8c, 0xc1, 0xc8, 0x49, 0x1d, 0x7d, 0x3b, 0x89, 0x19, 0xb8, 0x89, 0x89, 0xff, 0xf6, 0xe0,
	0xe9, 0x49, 0x9e, 0xab, 0x4d, 0xa5, 0xbf, 0xef, 0x3d, 0xbc, 0x81, 0x30, 0x57, 0xeb, 0xb5, 0xaa,
	0xd2, 0x9d, 0x4f, 0x8f, 0x76, 0xcd, 0x2d, 0x7f, 0xaf, 0x5c, 0x02, 0x6b, 0x6b, 0x91, 0xcb, 0x2b,
	0x99, 0x3b, 0xda, 0xc1, 0xeb, 0x21, 0x5a, 0xff, 0x3c, 0x79, 0x10, 0x37, 0x59, 0x74, 0xd2, 0x9e,
	0x78, 0x5b, 0xe9, 0x66, 0xcb, 0xc3, 0xf6, 0x01, 0xfd, 0xf2, 0x0c, 0x9e, 0x3f, 0x2a, 0x65, 0x21,
	0x0c, 0x6f, 0xc4, 0x96, 0xbc, 0x04, 0xdc, 0x7c, 0xb2, 0x67, 0x30, 0xbe, 0xcd, 0x56, 0x1b, 0xd1,
	0xdd, 0xca, 0x82, 0xef, 0x06, 0xdf, 0x78, 0xf1, 0x2f, 0x00, 0x5c, 0xe6, 0xe5, 0xff, 0x4f, 0x89,
	0x89, 0x59, 0xa9, 0x42, 0xb4, 0x58, 0x42, 0x43, 0x14, 0x5b, 0x60, 0x52, 0x2f, 0xb1, 0x9a, 0x1b,
	0x3c, 0x81, 0xea, 0x07, 0x1f, 0xad, 0xc7, 0xf1, 0x5f, 0x1e, 0xcc, 0x96, 0xbb, 0x32, 0x61, 0x73,
	0x18, 0x68, 0xd5, 0xd9, 0xc4, 0x2f, 0x1b, 0xb1, 0xca, 0xad, 0xcb, 0x80, 0x5b, 0x60, 0xd8, 0x5f,
	0x37, 0x4a, 0x67, 0x5d, 0xa9, 0x5a, 0xc0, 0x8e, 0xe0, 0x10, 0x43, 0xca, 0x22, 0xc5, 0xcc, 0xc9,
	0x55, 0x4a, 0xed, 0xd6, 0x1d, 0xf8, 0x94, 0x16, 0x2e, 0x0c, 0x4f, 0xdd, 0x62, 0x6e, 0x55, 0x64,
	0x18, 0xc0, 0x56, 0x2a, 0x7d, 0xef, 0x32, 0xe2, 0x3b, 0x19, 0x31, 0xe5, 0x48, 0x3d, 0x9b, 0x4a,
	0x5b, 0x8e, 0x4f, 0xb0, 0x39, 0x0d, 0xfe, 0xa1, 0x30, 0x6d, 0x7b, 0x2b, 0x9a, 0x16, 0x7d, 0x53,
	0x31, 0xe2, 0x4a, 0x07, 0xe3, 0xdf, 0x3d, 0x78, 0x7e, 0x51, 0x21, 0xc2, 0x97, 0x10, 0x85, 0x7b,
	0xc1, 0x04, 0x66, 0x4e, 0x5b, 0xd0, 0x4d, 0x67, 0xc7, 0x07, 0x89, 0x23, 0xe1, 0xae, 0xc0, 0x34,
	0x6d, 0x2b, 0xaf, 0x2b, 0x7c, 0x8a, 0xc6, 0x26, 0xc1, 0x74, 0x5f, 0x4f, 0xb0, 0x4f, 0xc0, 0xcf,
	0x9b, 0x6d, 0x8d, 0x29, 0x1b, 0xd2, 0x80, 0x98, 0x24, 0x67, 0x04, 0x79, 0x47, 0xc7, 0x7f, 0x78,
	0x70, 0xb8, 0x40, 0xf9, 0xbe, 0x89, 0x73, 0x78, 0xe6, 0x9c, 0x91, 0xfe, 0x26, 0x75, 0x99, 0x62,
	0xd0, 0xce, 0xcd, 0x8b, 0xe4, 0x51, 0xeb, 0x9c, 0x39, 0x7b, 0x7e, 0xc6, 0x2d, 0x18, 0x97, 0xbd,
	0x0f, 0x13, 0x7d, 0x97, 0x52, 0x81, 0x58, 0x73, 0xbe, 0xbe, 0x3b, 0xef, 0x4a, 0xc4, 0xd8, 0xc4,
	0x6e, 0x1b, 0x5a, 0xde, 0xa2, 0xf8, 0x0c, 0x02, 0x7a, 0x81, 0x53, 0x55, 0x6c, 0xd9, 0xd7, 0x70,
	0xe0, 0xce, 0x08, 0x3c, 0xdf, 0x94, 0x3f, 0x4b, 0xfe, 0xe5, 0x98, 0xef, 0xe9, 0xe2, 0x6b, 0x18,
	0xdb, 0x67, 0x74, 0x5e, 0xc0, 0xdb, 0x7b, 0x01, 0xf6, 0x99, 0x29, 0x51, 0x33, 0x12, 0xc9, 0x97,
	0x49, 0xb1, 0x33, 0x26, 0x79, 0xb7, 0x66, 0xc6, 0xeb, 0x25, 0x1a, 0x21, 0x8f, 0x33, 0x1c, 0xaf,
	0xf7, 0xd6, 0x38, 0xf1, 0xf1, 0x39, 0xcc, 0x89, 0x32, 0xd7, 0xed, 0x47, 0xf5, 0xf0, 0x72, 0x75,
	0xd3, 0x65, 0xca, 0xb7, 0x1b, 0xb8, 0xa1, 0x76, 0xf3, 0x67, 0xf0, 0xc8, 0xfc, 0x89, 0x4f, 0x61,
	0x4a, 0xda, 0xe5, 0x5d, 0xeb, 0xb4, 0x8f, 0xb7, 0xd7, 0x3e, 0xff, 0xe1, 0xe6, 0xe8, 0x5b, 0x08,
	0xee, 0xe7, 0x3f, 0x56, 0xf0, 0xfc, 0x64, 0xa3, 0x4b, 0xd5, 0x48, 0xbd, 0xe5, 0x38, 0x2c, 0x8a,
	0xf0, 0x3d, 0x36, 0x85, 0x11, 0xcf, 0xae, 0x74, 0xe8, 0x61, 0x1f, 0xc1, 0x52, 0x54, 0x78, 0xc5,
	0x35, 0xb6, 0x5a, 0x38, 0x38, 0x7a, 0x05, 0xbe, 0xad, 0x0c, 0xa3, 0x59, 0xbc, 0x3d, 0x7b, 0x87,
	0xea, 0x09, 0x0c, 0x17, 0x3f, 0x1d, 0x87, 0xde, 0xa5, 0x4f, 0xbf, 0xa3, 0x2f, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x27, 0x3b, 0x6e, 0xa2, 0x06, 0x00, 0x00,
}
