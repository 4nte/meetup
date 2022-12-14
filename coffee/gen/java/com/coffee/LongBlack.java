// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: grpcCoffe.proto

package com.coffee;

/**
 * Protobuf type {@code coffee.LongBlack}
 */
public final class LongBlack extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:coffee.LongBlack)
    LongBlackOrBuilder {
private static final long serialVersionUID = 0L;
  // Use LongBlack.newBuilder() to construct.
  private LongBlack(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private LongBlack() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new LongBlack();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.coffee.GrpcCoffeProto.internal_static_coffee_LongBlack_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.coffee.GrpcCoffeProto.internal_static_coffee_LongBlack_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.coffee.LongBlack.class, com.coffee.LongBlack.Builder.class);
  }

  public static final int SHOTS_OF_ESPRESSO_FIELD_NUMBER = 1;
  private int shotsOfEspresso_;
  /**
   * <pre>
   * 2 shots recommended
   * </pre>
   *
   * <code>int32 shots_of_espresso = 1 [json_name = "shotsOfEspresso"];</code>
   * @return The shotsOfEspresso.
   */
  @java.lang.Override
  public int getShotsOfEspresso() {
    return shotsOfEspresso_;
  }

  public static final int HOW_WATER_ML_FIELD_NUMBER = 2;
  private int howWaterMl_;
  /**
   * <pre>
   * 90 ml recommended
   * </pre>
   *
   * <code>int32 how_water_ml = 2 [json_name = "howWaterMl"];</code>
   * @return The howWaterMl.
   */
  @java.lang.Override
  public int getHowWaterMl() {
    return howWaterMl_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (shotsOfEspresso_ != 0) {
      output.writeInt32(1, shotsOfEspresso_);
    }
    if (howWaterMl_ != 0) {
      output.writeInt32(2, howWaterMl_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (shotsOfEspresso_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, shotsOfEspresso_);
    }
    if (howWaterMl_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, howWaterMl_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.coffee.LongBlack)) {
      return super.equals(obj);
    }
    com.coffee.LongBlack other = (com.coffee.LongBlack) obj;

    if (getShotsOfEspresso()
        != other.getShotsOfEspresso()) return false;
    if (getHowWaterMl()
        != other.getHowWaterMl()) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + SHOTS_OF_ESPRESSO_FIELD_NUMBER;
    hash = (53 * hash) + getShotsOfEspresso();
    hash = (37 * hash) + HOW_WATER_ML_FIELD_NUMBER;
    hash = (53 * hash) + getHowWaterMl();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.coffee.LongBlack parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.coffee.LongBlack parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.coffee.LongBlack parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.coffee.LongBlack parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.coffee.LongBlack parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.coffee.LongBlack parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.coffee.LongBlack parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.coffee.LongBlack parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.coffee.LongBlack parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.coffee.LongBlack parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.coffee.LongBlack parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.coffee.LongBlack parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.coffee.LongBlack prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code coffee.LongBlack}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:coffee.LongBlack)
      com.coffee.LongBlackOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.coffee.GrpcCoffeProto.internal_static_coffee_LongBlack_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.coffee.GrpcCoffeProto.internal_static_coffee_LongBlack_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.coffee.LongBlack.class, com.coffee.LongBlack.Builder.class);
    }

    // Construct using com.coffee.LongBlack.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      shotsOfEspresso_ = 0;

      howWaterMl_ = 0;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.coffee.GrpcCoffeProto.internal_static_coffee_LongBlack_descriptor;
    }

    @java.lang.Override
    public com.coffee.LongBlack getDefaultInstanceForType() {
      return com.coffee.LongBlack.getDefaultInstance();
    }

    @java.lang.Override
    public com.coffee.LongBlack build() {
      com.coffee.LongBlack result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.coffee.LongBlack buildPartial() {
      com.coffee.LongBlack result = new com.coffee.LongBlack(this);
      result.shotsOfEspresso_ = shotsOfEspresso_;
      result.howWaterMl_ = howWaterMl_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.coffee.LongBlack) {
        return mergeFrom((com.coffee.LongBlack)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.coffee.LongBlack other) {
      if (other == com.coffee.LongBlack.getDefaultInstance()) return this;
      if (other.getShotsOfEspresso() != 0) {
        setShotsOfEspresso(other.getShotsOfEspresso());
      }
      if (other.getHowWaterMl() != 0) {
        setHowWaterMl(other.getHowWaterMl());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {
              shotsOfEspresso_ = input.readInt32();

              break;
            } // case 8
            case 16: {
              howWaterMl_ = input.readInt32();

              break;
            } // case 16
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private int shotsOfEspresso_ ;
    /**
     * <pre>
     * 2 shots recommended
     * </pre>
     *
     * <code>int32 shots_of_espresso = 1 [json_name = "shotsOfEspresso"];</code>
     * @return The shotsOfEspresso.
     */
    @java.lang.Override
    public int getShotsOfEspresso() {
      return shotsOfEspresso_;
    }
    /**
     * <pre>
     * 2 shots recommended
     * </pre>
     *
     * <code>int32 shots_of_espresso = 1 [json_name = "shotsOfEspresso"];</code>
     * @param value The shotsOfEspresso to set.
     * @return This builder for chaining.
     */
    public Builder setShotsOfEspresso(int value) {
      
      shotsOfEspresso_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 2 shots recommended
     * </pre>
     *
     * <code>int32 shots_of_espresso = 1 [json_name = "shotsOfEspresso"];</code>
     * @return This builder for chaining.
     */
    public Builder clearShotsOfEspresso() {
      
      shotsOfEspresso_ = 0;
      onChanged();
      return this;
    }

    private int howWaterMl_ ;
    /**
     * <pre>
     * 90 ml recommended
     * </pre>
     *
     * <code>int32 how_water_ml = 2 [json_name = "howWaterMl"];</code>
     * @return The howWaterMl.
     */
    @java.lang.Override
    public int getHowWaterMl() {
      return howWaterMl_;
    }
    /**
     * <pre>
     * 90 ml recommended
     * </pre>
     *
     * <code>int32 how_water_ml = 2 [json_name = "howWaterMl"];</code>
     * @param value The howWaterMl to set.
     * @return This builder for chaining.
     */
    public Builder setHowWaterMl(int value) {
      
      howWaterMl_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 90 ml recommended
     * </pre>
     *
     * <code>int32 how_water_ml = 2 [json_name = "howWaterMl"];</code>
     * @return This builder for chaining.
     */
    public Builder clearHowWaterMl() {
      
      howWaterMl_ = 0;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:coffee.LongBlack)
  }

  // @@protoc_insertion_point(class_scope:coffee.LongBlack)
  private static final com.coffee.LongBlack DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.coffee.LongBlack();
  }

  public static com.coffee.LongBlack getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<LongBlack>
      PARSER = new com.google.protobuf.AbstractParser<LongBlack>() {
    @java.lang.Override
    public LongBlack parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<LongBlack> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<LongBlack> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.coffee.LongBlack getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

