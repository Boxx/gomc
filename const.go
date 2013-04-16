package gomc

/*
#include <libmemcached/memcached.h>
*/
import "C"

const (
	SUCCESS                          = C.MEMCACHED_SUCCESS
	FAILURE                          = C.MEMCACHED_FAILURE
	HOST_LOOKUP_FAILURE              = C.MEMCACHED_HOST_LOOKUP_FAILURE
	CONNECTION_FAILURE               = C.MEMCACHED_CONNECTION_FAILURE
	CONNECTION_BIND_FAILURE          = C.MEMCACHED_CONNECTION_BIND_FAILURE
	WRITE_FAILURE                    = C.MEMCACHED_WRITE_FAILURE
	READ_FAILURE                     = C.MEMCACHED_READ_FAILURE
	UNKNOWN_READ_FAILURE             = C.MEMCACHED_UNKNOWN_READ_FAILURE
	PROTOCOL_ERROR                   = C.MEMCACHED_PROTOCOL_ERROR
	CLIENT_ERROR                     = C.MEMCACHED_CLIENT_ERROR
	SERVER_ERROR                     = C.MEMCACHED_SERVER_ERROR
	CONNECTION_SOCKET_CREATE_FAILURE = C.MEMCACHED_CONNECTION_SOCKET_CREATE_FAILURE
	DATA_EXISTS                      = C.MEMCACHED_DATA_EXISTS
	DATA_DOES_NOT_EXIST              = C.MEMCACHED_DATA_DOES_NOT_EXIST
	NOTSTORED                        = C.MEMCACHED_NOTSTORED
	STORED                           = C.MEMCACHED_STORED
	NOTFOUND                         = C.MEMCACHED_NOTFOUND
	MEMORY_ALLOCATION_FAILURE        = C.MEMCACHED_MEMORY_ALLOCATION_FAILURE
	PARTIAL_READ                     = C.MEMCACHED_PARTIAL_READ
	SOME_ERRORS                      = C.MEMCACHED_SOME_ERRORS
	NO_SERVERS                       = C.MEMCACHED_NO_SERVERS
	END                              = C.MEMCACHED_END
	DELETED                          = C.MEMCACHED_DELETED
	VALUE                            = C.MEMCACHED_VALUE
	STAT                             = C.MEMCACHED_STAT
	ERRNO                            = C.MEMCACHED_ERRNO
	FAIL_UNIX_SOCKET                 = C.MEMCACHED_FAIL_UNIX_SOCKET
	NOT_SUPPORTED                    = C.MEMCACHED_NOT_SUPPORTED
	NO_KEY_PROVIDED                  = C.MEMCACHED_NO_KEY_PROVIDED
	FETCH_NOTFINISHED                = C.MEMCACHED_FETCH_NOTFINISHED
	TIMEOUT                          = C.MEMCACHED_TIMEOUT
	BUFFERED                         = C.MEMCACHED_BUFFERED
	BAD_KEY_PROVIDED                 = C.MEMCACHED_BAD_KEY_PROVIDED
	INVALID_HOST_PROTOCOL            = C.MEMCACHED_INVALID_HOST_PROTOCOL
	SERVER_MARKED_DEAD               = C.MEMCACHED_SERVER_MARKED_DEAD
	UNKNOWN_STAT_KEY                 = C.MEMCACHED_UNKNOWN_STAT_KEY
	E2BIG                            = C.MEMCACHED_E2BIG
	INVALID_ARGUMENTS                = C.MEMCACHED_INVALID_ARGUMENTS
	KEY_TOO_BIG                      = C.MEMCACHED_KEY_TOO_BIG
	AUTH_PROBLEM                     = C.MEMCACHED_AUTH_PROBLEM
	AUTH_FAILURE                     = C.MEMCACHED_AUTH_FAILURE
	AUTH_CONTINUE                    = C.MEMCACHED_AUTH_CONTINUE
	PARSE_ERROR                      = C.MEMCACHED_PARSE_ERROR
	PARSE_USER_ERROR                 = C.MEMCACHED_PARSE_USER_ERROR
	DEPRECATED                       = C.MEMCACHED_DEPRECATED
	IN_PROGRESS                      = C.MEMCACHED_IN_PROGRESS
	SERVER_TEMPORARILY_DISABLED      = C.MEMCACHED_SERVER_TEMPORARILY_DISABLED
	SERVER_MEMORY_ALLOCATION_FAILURE = C.MEMCACHED_SERVER_MEMORY_ALLOCATION_FAILURE
	MAXIMUM_RETURN                   = C.MEMCACHED_MAXIMUM_RETURN

	BEHAVIOR_NO_BLOCK               = C.MEMCACHED_BEHAVIOR_NO_BLOCK
	BEHAVIOR_TCP_NODELAY            = C.MEMCACHED_BEHAVIOR_TCP_NODELAY
	BEHAVIOR_HASH                   = C.MEMCACHED_BEHAVIOR_HASH
	BEHAVIOR_KETAMA                 = C.MEMCACHED_BEHAVIOR_KETAMA
	BEHAVIOR_SOCKET_SEND_SIZE       = C.MEMCACHED_BEHAVIOR_SOCKET_SEND_SIZE
	BEHAVIOR_SOCKET_RECV_SIZE       = C.MEMCACHED_BEHAVIOR_SOCKET_RECV_SIZE
	BEHAVIOR_CACHE_LOOKUPS          = C.MEMCACHED_BEHAVIOR_CACHE_LOOKUPS
	BEHAVIOR_SUPPORT_CAS            = C.MEMCACHED_BEHAVIOR_SUPPORT_CAS
	BEHAVIOR_POLL_TIMEOUT           = C.MEMCACHED_BEHAVIOR_POLL_TIMEOUT
	BEHAVIOR_DISTRIBUTION           = C.MEMCACHED_BEHAVIOR_DISTRIBUTION
	BEHAVIOR_BUFFER_REQUESTS        = C.MEMCACHED_BEHAVIOR_BUFFER_REQUESTS
	BEHAVIOR_USER_DATA              = C.MEMCACHED_BEHAVIOR_USER_DATA
	BEHAVIOR_SORT_HOSTS             = C.MEMCACHED_BEHAVIOR_SORT_HOSTS
	BEHAVIOR_VERIFY_KEY             = C.MEMCACHED_BEHAVIOR_VERIFY_KEY
	BEHAVIOR_CONNECT_TIMEOUT        = C.MEMCACHED_BEHAVIOR_CONNECT_TIMEOUT
	BEHAVIOR_RETRY_TIMEOUT          = C.MEMCACHED_BEHAVIOR_RETRY_TIMEOUT
	BEHAVIOR_KETAMA_WEIGHTED        = C.MEMCACHED_BEHAVIOR_KETAMA_WEIGHTED
	BEHAVIOR_KETAMA_HASH            = C.MEMCACHED_BEHAVIOR_KETAMA_HASH
	BEHAVIOR_BINARY_PROTOCOL        = C.MEMCACHED_BEHAVIOR_BINARY_PROTOCOL
	BEHAVIOR_SND_TIMEOUT            = C.MEMCACHED_BEHAVIOR_SND_TIMEOUT
	BEHAVIOR_RCV_TIMEOUT            = C.MEMCACHED_BEHAVIOR_RCV_TIMEOUT
	BEHAVIOR_SERVER_FAILURE_LIMIT   = C.MEMCACHED_BEHAVIOR_SERVER_FAILURE_LIMIT
	BEHAVIOR_IO_MSG_WATERMARK       = C.MEMCACHED_BEHAVIOR_IO_MSG_WATERMARK
	BEHAVIOR_IO_BYTES_WATERMARK     = C.MEMCACHED_BEHAVIOR_IO_BYTES_WATERMARK
	BEHAVIOR_IO_KEY_PREFETCH        = C.MEMCACHED_BEHAVIOR_IO_KEY_PREFETCH
	BEHAVIOR_HASH_WITH_PREFIX_KEY   = C.MEMCACHED_BEHAVIOR_HASH_WITH_PREFIX_KEY
	BEHAVIOR_NOREPLY                = C.MEMCACHED_BEHAVIOR_NOREPLY
	BEHAVIOR_USE_UDP                = C.MEMCACHED_BEHAVIOR_USE_UDP
	BEHAVIOR_AUTO_EJECT_HOSTS       = C.MEMCACHED_BEHAVIOR_AUTO_EJECT_HOSTS
	BEHAVIOR_NUMBER_OF_REPLICAS     = C.MEMCACHED_BEHAVIOR_NUMBER_OF_REPLICAS
	BEHAVIOR_RANDOMIZE_REPLICA_READ = C.MEMCACHED_BEHAVIOR_RANDOMIZE_REPLICA_READ
	BEHAVIOR_CORK                   = C.MEMCACHED_BEHAVIOR_CORK
	BEHAVIOR_TCP_KEEPALIVE          = C.MEMCACHED_BEHAVIOR_TCP_KEEPALIVE
	BEHAVIOR_TCP_KEEPIDLE           = C.MEMCACHED_BEHAVIOR_TCP_KEEPIDLE
	BEHAVIOR_LOAD_FROM_FILE         = C.MEMCACHED_BEHAVIOR_LOAD_FROM_FILE
	BEHAVIOR_REMOVE_FAILED_SERVERS  = C.MEMCACHED_BEHAVIOR_REMOVE_FAILED_SERVERS
	BEHAVIOR_DEAD_TIMEOUT           = C.MEMCACHED_BEHAVIOR_DEAD_TIMEOUT
	BEHAVIOR_MAX                    = C.MEMCACHED_BEHAVIOR_MAX

	DISTRIBUTION_MODULA                = C.MEMCACHED_DISTRIBUTION_MODULA
	DISTRIBUTION_CONSISTENT            = C.MEMCACHED_DISTRIBUTION_CONSISTENT
	DISTRIBUTION_CONSISTENT_KETAMA     = C.MEMCACHED_DISTRIBUTION_CONSISTENT_KETAMA
	DISTRIBUTION_RANDOM                = C.MEMCACHED_DISTRIBUTION_RANDOM
	DISTRIBUTION_CONSISTENT_KETAMA_SPY = C.MEMCACHED_DISTRIBUTION_CONSISTENT_KETAMA_SPY
	DISTRIBUTION_CONSISTENT_WEIGHTED   = C.MEMCACHED_DISTRIBUTION_CONSISTENT_WEIGHTED
	DISTRIBUTION_VIRTUAL_BUCKET        = C.MEMCACHED_DISTRIBUTION_VIRTUAL_BUCKET
	DISTRIBUTION_CONSISTENT_MAX        = C.MEMCACHED_DISTRIBUTION_CONSISTENT_MAX

	HASH_DEFAULT  = C.MEMCACHED_HASH_DEFAULT
	HASH_MD5      = C.MEMCACHED_HASH_MD5
	HASH_CRC      = C.MEMCACHED_HASH_CRC
	HASH_FNV1_64  = C.MEMCACHED_HASH_FNV1_64
	HASH_FNV1A_64 = C.MEMCACHED_HASH_FNV1A_64
	HASH_FNV1_32  = C.MEMCACHED_HASH_FNV1_32
	HASH_FNV1A_32 = C.MEMCACHED_HASH_FNV1A_32
	HASH_HSIEH    = C.MEMCACHED_HASH_HSIEH
	HASH_MURMUR   = C.MEMCACHED_HASH_MURMUR
	HASH_JENKINS  = C.MEMCACHED_HASH_JENKINS
	HASH_CUSTOM   = C.MEMCACHED_HASH_CUSTOM
	HASH_MAX      = C.MEMCACHED_HASH_MAX

	CONNECTION_TCP         = C.MEMCACHED_CONNECTION_TCP
	CONNECTION_UDP         = C.MEMCACHED_CONNECTION_UDP
	CONNECTION_UNIX_SOCKET = C.MEMCACHED_CONNECTION_UNIX_SOCKET

	DEFAULT_PORT = C.MEMCACHED_DEFAULT_PORT
)