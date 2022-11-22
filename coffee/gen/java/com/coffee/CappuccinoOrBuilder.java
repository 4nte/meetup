// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: grpcCoffe.proto

package com.coffee;

public interface CappuccinoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:coffee.Cappuccino)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * 1-2 shots of espresso
   * </pre>
   *
   * <code>int32 shots_of_espresso = 1 [json_name = "shotsOfEspresso"];</code>
   * @return The shotsOfEspresso.
   */
  int getShotsOfEspresso();

  /**
   * <pre>
   * 60 ml recommended
   * </pre>
   *
   * <code>int32 steamed_milk_ml = 2 [json_name = "steamedMilkMl"];</code>
   * @return The steamedMilkMl.
   */
  int getSteamedMilkMl();

  /**
   * <pre>
   * 60 ml recommended
   * </pre>
   *
   * <code>int32 foamed_milk_ml = 3 [json_name = "foamedMilkMl"];</code>
   * @return The foamedMilkMl.
   */
  int getFoamedMilkMl();

  /**
   * <pre>
   * optional sprinkling of chocolate powder
   * </pre>
   *
   * <code>bool sprinklinkg_of_chocolate_poweder = 4 [json_name = "sprinklinkgOfChocolatePoweder"];</code>
   * @return The sprinklinkgOfChocolatePoweder.
   */
  boolean getSprinklinkgOfChocolatePoweder();
}
