// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: coffe.proto

package com.coffe;

public final class CoffeProto {
  private CoffeProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_coffe_MakeCoffeRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_coffe_MakeCoffeRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\013coffe.proto\022\005coffe\"8\n\020MakeCoffeRequest" +
      "\022$\n\004type\030\001 \001(\0162\020.coffe.CoffeTypeR\004type*t" +
      "\n\tCoffeType\022\032\n\026COFFE_TYPE_UNSPECIFIED\020\000\022" +
      "\027\n\023COFFE_TYPE_ESPRESSO\020\001\022\030\n\024COFFE_TYPE_C" +
      "APUCCINO\020\002\022\030\n\024COFFE_TYPE_AMERICANO\020\003Bk\n\t" +
      "com.coffeB\nCoffeProtoP\001Z\033github.com/4nte" +
      "/coffe/proto\370\001\000\242\002\003CXX\252\002\005Coffe\312\002\005Coffe\342\002\021" +
      "Coffe\\GPBMetadata\352\002\005Coffeb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_coffe_MakeCoffeRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_coffe_MakeCoffeRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_coffe_MakeCoffeRequest_descriptor,
        new java.lang.String[] { "Type", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
