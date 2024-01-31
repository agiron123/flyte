// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/cache.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fservice_2fcache_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fservice_2fcache_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "google/api/annotations.pb.h"
#include "flyteidl/core/errors.pb.h"
#include "flyteidl/core/identifier.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fservice_2fcache_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fservice_2fcache_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fservice_2fcache_2eproto();
namespace flyteidl {
namespace service {
class EvictCacheResponse;
class EvictCacheResponseDefaultTypeInternal;
extern EvictCacheResponseDefaultTypeInternal _EvictCacheResponse_default_instance_;
class EvictTaskExecutionCacheRequest;
class EvictTaskExecutionCacheRequestDefaultTypeInternal;
extern EvictTaskExecutionCacheRequestDefaultTypeInternal _EvictTaskExecutionCacheRequest_default_instance_;
}  // namespace service
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::service::EvictCacheResponse* Arena::CreateMaybeMessage<::flyteidl::service::EvictCacheResponse>(Arena*);
template<> ::flyteidl::service::EvictTaskExecutionCacheRequest* Arena::CreateMaybeMessage<::flyteidl::service::EvictTaskExecutionCacheRequest>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace service {

// ===================================================================

class EvictTaskExecutionCacheRequest final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.service.EvictTaskExecutionCacheRequest) */ {
 public:
  EvictTaskExecutionCacheRequest();
  virtual ~EvictTaskExecutionCacheRequest();

  EvictTaskExecutionCacheRequest(const EvictTaskExecutionCacheRequest& from);

  inline EvictTaskExecutionCacheRequest& operator=(const EvictTaskExecutionCacheRequest& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  EvictTaskExecutionCacheRequest(EvictTaskExecutionCacheRequest&& from) noexcept
    : EvictTaskExecutionCacheRequest() {
    *this = ::std::move(from);
  }

  inline EvictTaskExecutionCacheRequest& operator=(EvictTaskExecutionCacheRequest&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const EvictTaskExecutionCacheRequest& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const EvictTaskExecutionCacheRequest* internal_default_instance() {
    return reinterpret_cast<const EvictTaskExecutionCacheRequest*>(
               &_EvictTaskExecutionCacheRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(EvictTaskExecutionCacheRequest* other);
  friend void swap(EvictTaskExecutionCacheRequest& a, EvictTaskExecutionCacheRequest& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline EvictTaskExecutionCacheRequest* New() const final {
    return CreateMaybeMessage<EvictTaskExecutionCacheRequest>(nullptr);
  }

  EvictTaskExecutionCacheRequest* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<EvictTaskExecutionCacheRequest>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const EvictTaskExecutionCacheRequest& from);
  void MergeFrom(const EvictTaskExecutionCacheRequest& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(EvictTaskExecutionCacheRequest* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.core.TaskExecutionIdentifier task_execution_id = 1;
  bool has_task_execution_id() const;
  void clear_task_execution_id();
  static const int kTaskExecutionIdFieldNumber = 1;
  const ::flyteidl::core::TaskExecutionIdentifier& task_execution_id() const;
  ::flyteidl::core::TaskExecutionIdentifier* release_task_execution_id();
  ::flyteidl::core::TaskExecutionIdentifier* mutable_task_execution_id();
  void set_allocated_task_execution_id(::flyteidl::core::TaskExecutionIdentifier* task_execution_id);

  // @@protoc_insertion_point(class_scope:flyteidl.service.EvictTaskExecutionCacheRequest)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::flyteidl::core::TaskExecutionIdentifier* task_execution_id_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fservice_2fcache_2eproto;
};
// -------------------------------------------------------------------

class EvictCacheResponse final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.service.EvictCacheResponse) */ {
 public:
  EvictCacheResponse();
  virtual ~EvictCacheResponse();

  EvictCacheResponse(const EvictCacheResponse& from);

  inline EvictCacheResponse& operator=(const EvictCacheResponse& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  EvictCacheResponse(EvictCacheResponse&& from) noexcept
    : EvictCacheResponse() {
    *this = ::std::move(from);
  }

  inline EvictCacheResponse& operator=(EvictCacheResponse&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const EvictCacheResponse& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const EvictCacheResponse* internal_default_instance() {
    return reinterpret_cast<const EvictCacheResponse*>(
               &_EvictCacheResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(EvictCacheResponse* other);
  friend void swap(EvictCacheResponse& a, EvictCacheResponse& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline EvictCacheResponse* New() const final {
    return CreateMaybeMessage<EvictCacheResponse>(nullptr);
  }

  EvictCacheResponse* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<EvictCacheResponse>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const EvictCacheResponse& from);
  void MergeFrom(const EvictCacheResponse& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(EvictCacheResponse* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.core.CacheEvictionErrorList errors = 1;
  bool has_errors() const;
  void clear_errors();
  static const int kErrorsFieldNumber = 1;
  const ::flyteidl::core::CacheEvictionErrorList& errors() const;
  ::flyteidl::core::CacheEvictionErrorList* release_errors();
  ::flyteidl::core::CacheEvictionErrorList* mutable_errors();
  void set_allocated_errors(::flyteidl::core::CacheEvictionErrorList* errors);

  // @@protoc_insertion_point(class_scope:flyteidl.service.EvictCacheResponse)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::flyteidl::core::CacheEvictionErrorList* errors_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fservice_2fcache_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// EvictTaskExecutionCacheRequest

// .flyteidl.core.TaskExecutionIdentifier task_execution_id = 1;
inline bool EvictTaskExecutionCacheRequest::has_task_execution_id() const {
  return this != internal_default_instance() && task_execution_id_ != nullptr;
}
inline const ::flyteidl::core::TaskExecutionIdentifier& EvictTaskExecutionCacheRequest::task_execution_id() const {
  const ::flyteidl::core::TaskExecutionIdentifier* p = task_execution_id_;
  // @@protoc_insertion_point(field_get:flyteidl.service.EvictTaskExecutionCacheRequest.task_execution_id)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::core::TaskExecutionIdentifier*>(
      &::flyteidl::core::_TaskExecutionIdentifier_default_instance_);
}
inline ::flyteidl::core::TaskExecutionIdentifier* EvictTaskExecutionCacheRequest::release_task_execution_id() {
  // @@protoc_insertion_point(field_release:flyteidl.service.EvictTaskExecutionCacheRequest.task_execution_id)
  
  ::flyteidl::core::TaskExecutionIdentifier* temp = task_execution_id_;
  task_execution_id_ = nullptr;
  return temp;
}
inline ::flyteidl::core::TaskExecutionIdentifier* EvictTaskExecutionCacheRequest::mutable_task_execution_id() {
  
  if (task_execution_id_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::core::TaskExecutionIdentifier>(GetArenaNoVirtual());
    task_execution_id_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.service.EvictTaskExecutionCacheRequest.task_execution_id)
  return task_execution_id_;
}
inline void EvictTaskExecutionCacheRequest::set_allocated_task_execution_id(::flyteidl::core::TaskExecutionIdentifier* task_execution_id) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(task_execution_id_);
  }
  if (task_execution_id) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      task_execution_id = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, task_execution_id, submessage_arena);
    }
    
  } else {
    
  }
  task_execution_id_ = task_execution_id;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.service.EvictTaskExecutionCacheRequest.task_execution_id)
}

// -------------------------------------------------------------------

// EvictCacheResponse

// .flyteidl.core.CacheEvictionErrorList errors = 1;
inline bool EvictCacheResponse::has_errors() const {
  return this != internal_default_instance() && errors_ != nullptr;
}
inline const ::flyteidl::core::CacheEvictionErrorList& EvictCacheResponse::errors() const {
  const ::flyteidl::core::CacheEvictionErrorList* p = errors_;
  // @@protoc_insertion_point(field_get:flyteidl.service.EvictCacheResponse.errors)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::core::CacheEvictionErrorList*>(
      &::flyteidl::core::_CacheEvictionErrorList_default_instance_);
}
inline ::flyteidl::core::CacheEvictionErrorList* EvictCacheResponse::release_errors() {
  // @@protoc_insertion_point(field_release:flyteidl.service.EvictCacheResponse.errors)
  
  ::flyteidl::core::CacheEvictionErrorList* temp = errors_;
  errors_ = nullptr;
  return temp;
}
inline ::flyteidl::core::CacheEvictionErrorList* EvictCacheResponse::mutable_errors() {
  
  if (errors_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::core::CacheEvictionErrorList>(GetArenaNoVirtual());
    errors_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.service.EvictCacheResponse.errors)
  return errors_;
}
inline void EvictCacheResponse::set_allocated_errors(::flyteidl::core::CacheEvictionErrorList* errors) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(errors_);
  }
  if (errors) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      errors = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, errors, submessage_arena);
    }
    
  } else {
    
  }
  errors_ = errors;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.service.EvictCacheResponse.errors)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace service
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fservice_2fcache_2eproto